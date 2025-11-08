package model

type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at"` // could be time.Time
	UpdatedAt string   `json:"updated_at"`
}
