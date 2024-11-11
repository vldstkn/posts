package app

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"p1/internal/config"
	"p1/internal/repository/postgres"
	"p1/internal/services"
	transport "p1/internal/transport/http/handlers"
	"p1/internal/transport/http/middleware"
	"p1/pkg/db"
)

func App(logger *slog.Logger, conf *config.Config) http.Handler {
	dbPg := db.NewDB(conf.DBString)

	router := chi.NewRouter()
	router.Use(middleware.Logger(logger))

	//Repositories
	userRepository := postgres.NewUserRepository(dbPg)
	postRepository := postgres.NewPostService(dbPg)

	//Services
	userService := services.NewUserService(services.UserServiceDeps{
		UserRepository: userRepository,
	})
	authService := services.NewAuthService(services.AuthServiceDeps{
		UserService: userService,
	})
	postService := services.NewPostService(services.PostServiceDeps{
		PostRepository: postRepository,
	})

	//Handlers
	transport.NewAuthHandler(router, transport.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
		Logger:      logger,
	})
	transport.NewPostHandler(router, transport.PostHandlerDeps{
		PostService: postService,
		Logger:      logger,
		Config:      conf,
	})
	return router
}
