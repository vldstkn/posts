package services

import (
	"p1/internal/di"
	"p1/internal/domain"
	"strconv"
)

type PostServiceDeps struct {
	PostRepository di.IPostRepository
}

type PostService struct {
	PostRepository di.IPostRepository
}

func NewPostService(deps PostServiceDeps) *PostService {
	return &PostService{
		PostRepository: deps.PostRepository,
	}
}

func (service *PostService) Create(title, content string, authorId int) (int, error) {
	post := &domain.Post{
		Content:  content,
		Title:    title,
		AuthorId: authorId,
	}
	postId, err := service.PostRepository.Create(post)
	if err != nil {
		return -1, err
	}
	return postId, nil
}

func (service *PostService) FindById(id string) (*domain.Post, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	post := service.PostRepository.FindById(idInt)
	return post, nil
}

func (service *PostService) DeleteById(id string) error {
	return nil
}

func (service *PostService) FindByTitle(title string) []domain.Post {
	return []domain.Post{}
}
