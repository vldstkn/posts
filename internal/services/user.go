package services

import (
	"errors"
	"p1/internal/di"
	"p1/internal/domain"
)

type UserServiceDeps struct {
	UserRepository di.IUserRepository
}

type UserService struct {
	UserRepository di.IUserRepository
}

func NewUserService(deps UserServiceDeps) *UserService {
	return &UserService{
		UserRepository: deps.UserRepository,
	}
}

func (service *UserService) Create(email, password string) (int, error) {
	existsUser := service.UserRepository.FindByEmail(email)
	if existsUser != nil {
		return -1, errors.New("user exists")
	}
	user := &domain.User{
		Email:    email,
		Password: password,
	}
	userId, err := service.UserRepository.Create(user)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (service *UserService) GetByEmail(email string) *domain.User {
	user := service.UserRepository.FindByEmail(email)
	return user
}
