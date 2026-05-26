package subscriptions

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

func (r *Repository) Create(ctx context.Context, input CreateSubscriptionInput) (Subscription, error) {
	query := `
		INSERT INTO subscriptions (user_id, flight_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, flight_id)
		DO UPDATE SET active = true
		RETURNING user_id::text, flight_id, active, created_at
	`

	var subscription Subscription
	err := r.db.QueryRowContext(ctx, query, input.UserID, input.FlightID).Scan(
		&subscription.UserID,
		&subscription.FlightID,
		&subscription.Active,
		&subscription.CreatedAt,
	)
	if err != nil {
		return Subscription{}, err
	}

	return subscription, nil
}
