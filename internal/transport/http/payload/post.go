package payload

import "p1/internal/domain"

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CreatePostResponse struct {
	Id int `json:"id"`
}

type FindByIdPostRequest struct {
	Post domain.Post `json:"post"`
}
