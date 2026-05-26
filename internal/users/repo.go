package users

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, input CreateUserInput) (User, error) {
	query := `
		INSERT INTO users (name, number)
		VALUES ($1, $2)
		RETURNING id::text, name, number, created_at
	`

	var user User
	err := r.db.QueryRowContext(ctx, query, input.Name, input.PhoneNumber).Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
