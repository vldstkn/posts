package di

import (
	"net/http"
	"p1/internal/domain"
	"p1/pkg/jwt"
	"time"
)

type IUserService interface {
	Create(email, password string) error
}

type IAuthService interface {
	Register(email, password string) (int, error)
	Login(email, password string) (int, string, error)
	IssueTokens(data jwt.JWTData, secret string, expirationTime time.Time) (string, string, error)
	AddCookie(w http.ResponseWriter, name, value string, expirationTime time.Time)
	GetNewTokens(refreshToken, secret string, expirationTime time.Time) (string, string, error)
}
type IPostService interface {
	Create(title, content string, authorId int) (int, error)
	DeleteById(id string) error
	FindById(id string) (*domain.Post, error)
	FindByTitle(title string) []domain.Post
}
