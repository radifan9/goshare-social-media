package model

import "time"

type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
