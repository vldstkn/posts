package domain

import "time"

type Post struct {
	Id        int       `json:"id" db:"id""`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	AuthorId  int       `json:"author_id" db:"author_id"`
}
