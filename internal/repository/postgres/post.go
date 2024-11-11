package postgres

import (
	"p1/internal/domain"
	"p1/pkg/db"
)

type PostRepository struct {
	*db.DB
}

func NewPostService(database *db.DB) *PostRepository {
	return &PostRepository{
		database,
	}
}

func (repo *PostRepository) Create(post *domain.Post) (int, error) {
	var id int
	err := repo.DB.QueryRow("INSERT INTO post (title, content, author_id) VALUES ($1, $2, $3) RETURNING id",
		post.Title, post.Content, post.AuthorId).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repo *PostRepository) FindById(id int) *domain.Post {
	post := domain.Post{}

	err := repo.DB.Get(&post, "SELECT * FROM Post WHERE id=$1", id)

	if err != nil {
		return nil
	}
	return &post
}

func (repo *PostRepository) DeleteById(id string) error {
	return nil
}

func (repo *PostRepository) FindByTitle(title string) []domain.Post {
	return []domain.Post{}
}
