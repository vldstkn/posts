package handlers

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"p1/internal/config"
	"p1/internal/di"
	server_error "p1/internal/transport/http/error"
	"p1/internal/transport/http/middleware"
	"p1/internal/transport/http/payload"
	"p1/pkg/req"
	"p1/pkg/res"
)

type PostHandlerDeps struct {
	PostService di.IPostService
	Logger      *slog.Logger
	*config.Config
}

type PostHandler struct {
	PostService di.IPostService
	Logger      *slog.Logger
	*config.Config
}

func NewPostHandler(router *chi.Mux, deps PostHandlerDeps) {
	handler := PostHandler{
		PostService: deps.PostService,
		Logger:      deps.Logger,
		Config:      deps.Config,
	}
	router.Route("/post", func(router chi.Router) {
		router.Use(middleware.IsAuthed(handler.Config))
		router.Group(func(router chi.Router) {
			router.Use(middleware.IsAuthor(handler.JWTSecret))
			router.Post("/", handler.Create())
		})
		router.Get("/{id}", handler.FindById())
	})
}

func (handler *PostHandler) Create() http.HandlerFunc {
	opt := "PostHandler.Create"
	return func(w http.ResponseWriter, r *http.Request) {
		authInfo := r.Context().Value("authInfo").(middleware.AuthInfo)
		body, err := req.HandleBody[payload.CreatePostRequest](&w, r)
		if err != nil {
			server_error.InternalServerError(w, handler.Logger, opt, err)
			return
		}
		postId, err := handler.PostService.Create(body.Title, body.Content, authInfo.Id)
		if err != nil {
			server_error.InternalServerError(w, handler.Logger, opt, err)
			return
		}

		res.Json(w, payload.CreatePostResponse{
			Id: postId,
		}, http.StatusCreated)
	}
}

func (handler *PostHandler) FindById() http.HandlerFunc {
	opt := "PostHandler.FindById"
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		post, err := handler.PostService.FindById(id)
		if err != nil {
			server_error.BadRequest(w, handler.Logger, opt, err)
			return
		}
		if post == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		res.Json(w, payload.FindByIdPostRequest{
			Post: *post,
		}, http.StatusOK)
	}
}
