package postgres

import (
	"p1/internal/domain"
	"p1/pkg/db"
)

type UserRepository struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user *domain.User) (int, error) {
	var id int
	err := repo.db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", user.Email, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repo *UserRepository) FindByEmail(email string) *domain.User {
	var user domain.User
	err := repo.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil
	}
	return &user
}

func (repo *UserRepository) FindById(test string) *domain.User {
	return nil
}

func (repo *UserRepository) DeleteById(id string) bool {
	return false
}
