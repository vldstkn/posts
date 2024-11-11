package handlers

import (
	"log/slog"
	"net/http"
	"p1/internal/config"
	"p1/internal/di"
	server_error "p1/internal/transport/http/error"
	"p1/internal/transport/http/payload"
	"p1/pkg/jwt"
	"p1/pkg/req"
	"p1/pkg/res"
	"time"

	"github.com/go-chi/chi/v5"
)

type AuthHandlerDeps struct {
	*config.Config
	AuthService di.IAuthService
	Logger      *slog.Logger
}

type AuthHandler struct {
	*config.Config
	AuthService di.IAuthService
	Logger      *slog.Logger
}

func NewAuthHandler(router *chi.Mux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
		Logger:      deps.Logger,
		Config:      deps.Config,
	}
	router.Route("/auth", func(router chi.Router) {
		router.Post("/register", handler.Register())
		router.Post("/login", handler.Login())
		router.Get("/login/access-token", handler.GetNewTokens())
	})

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := "AuthHandler.Register"
		body, err := req.HandleBody[payload.RegisterRequest](&w, r)
		if err != nil {
			server_error.InternalServerError(w, handler.Logger, opt, err)
			return
		}
		userId, err := handler.AuthService.Register(body.Email, body.Password)
		if err != nil {
			server_error.InternalServerError(w, handler.Logger, opt, err)
			return
		}
		expirationTime := time.Now().Add(time.Hour)
		accessToken, refreshToken, err := handler.AuthService.IssueTokens(jwt.JWTData{
			Id:   userId,
			Role: "user",
		}, handler.Config.JWTSecret, expirationTime)
		if err != nil {
			server_error.InternalServerError(w, handler.Logger, opt, err)
		}
		handler.AuthService.AddCookie(w, "refresh_token", refreshToken, expirationTime)
		res.Json(w, payload.RegisterResponse{
			Id:          userId,
			AccessToken: accessToken,
		}, 201)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := "AuthHandler.Login"
		body, err := req.HandleBody[payload.LoginRequest](&w, r)
		if err != nil {
			server_error.BadRequest(w, handler.Logger, opt, err)
			return
		}
		userId, userRole, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, "invalid email or password", http.StatusBadRequest)
			return
		}
		expirationTime := time.Now().Add(time.Hour)
		accessToken, refreshToken, err := handler.AuthService.IssueTokens(jwt.JWTData{
			Id:   userId,
			Role: userRole,
		}, handler.Config.JWTSecret, expirationTime)

		if err != nil {
			server_error.BadRequest(w, handler.Logger, opt, err)
			return
		}

		handler.AuthService.AddCookie(w, "refresh_token", refreshToken, expirationTime)
		res.Json(w, payload.LoginResponse{
			Id:          userId,
			AccessToken: accessToken,
		}, http.StatusOK)
	}
}

func (handler *AuthHandler) GetNewTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenFromCookie, err := r.Cookie("refresh_token")
		if tokenFromCookie.Value == "" || err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		expirationTime := time.Now().Add(time.Hour)
		accessToken, refreshToken, err := handler.AuthService.GetNewTokens(tokenFromCookie.Value, handler.JWTSecret, expirationTime)
		if err != nil {
			// TODO:
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		handler.AuthService.AddCookie(w, "refresh_token", refreshToken, expirationTime.Add(time.Hour*3))
		res.Json(w, payload.GetNewTokensResponse{
			AccessToken: accessToken,
		}, http.StatusOK)
	}
}
