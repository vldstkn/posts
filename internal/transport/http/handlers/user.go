package handlers

import (
	"github.com/go-chi/chi/v5"
	"p1/internal/di"
)

type UserHandlerDeps struct {
	UserService di.IUserService
}

func NewUserHandler(router *chi.Mux, deps UserHandlerDeps) {

}
