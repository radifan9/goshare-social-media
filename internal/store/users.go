package store

import (
	"context"
	"database/sql"
	"social/internal/model"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}
