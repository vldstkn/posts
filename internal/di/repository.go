package di

import "p1/internal/domain"

type IUserRepository interface {
	Create(user *domain.User) (int, error)
	FindByEmail(email string) *domain.User
}

type IPostRepository interface {
	Create(post *domain.Post) (int, error)
	DeleteById(id string) error
	FindById(id int) *domain.Post
	FindByTitle(title string) []domain.Post
}
