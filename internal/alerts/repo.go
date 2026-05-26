package alerts

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

func (r *Repository) CreateForSubscribers(ctx context.Context, flightID, alertType, eventKey, message string) ([]Alert, error) {
	query := `
		INSERT INTO alerts (user_id, flight_id, alert_type, message, event_key)
		SELECT user_id, flight_id, $2, $3, $4
		FROM subscriptions
		WHERE flight_id = $1 AND active = true
		ON CONFLICT (user_id, flight_id, event_key) DO NOTHING
		RETURNING id, user_id::text, flight_id, alert_type, message, event_key, created_at
	`

	rows, err := r.db.QueryContext(ctx, query, flightID, alertType, message, eventKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var created []Alert
	for rows.Next() {
		var alert Alert
		if err := rows.Scan(
			&alert.ID,
			&alert.UserID,
			&alert.FlightID,
			&alert.Type,
			&alert.Message,
			&alert.EventKey,
			&alert.CreatedAt,
		); err != nil {
			return nil, err
		}
		created = append(created, alert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return created, nil
}

func (r *Repository) MarkSent(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `UPDATE alerts SET sent_at = now() WHERE id = $1`, id)
	return err
}
