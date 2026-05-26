package flights

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) EnsureFlight(ctx context.Context, update FlightUpdate) error {
	query := `
		INSERT INTO flights (
			flight_id,
			flight_number,
			status,
			departure_gate,
			departure_terminal,
			delay_minutes,
			last_updated
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (flight_id) DO UPDATE SET
			flight_number = EXCLUDED.flight_number,
			status = EXCLUDED.status,
			departure_gate = EXCLUDED.departure_gate,
			departure_terminal = EXCLUDED.departure_terminal,
			delay_minutes = EXCLUDED.delay_minutes,
			last_updated = EXCLUDED.last_updated
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		update.FlightID,
		update.FlightNumber,
		update.Status,
		update.DepartureGate,
		update.DepartureTerminal,
		update.DelayMinutes,
		update.ObservedAt,
	)
	return err
}

func (r *Repository) GetState(ctx context.Context, flightID string) (FlightState, bool, error) {
	query := `
		SELECT flight_id, status, departure_gate, departure_terminal, delay_minutes, updated_at
		FROM flight_state
		WHERE flight_id = $1
	`

	var state FlightState
	err := r.db.QueryRowContext(ctx, query, flightID).Scan(
		&state.FlightID,
		&state.Status,
		&state.DepartureGate,
		&state.DepartureTerminal,
		&state.DelayMinutes,
		&state.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return FlightState{}, false, nil
	}
	if err != nil {
		return FlightState{}, false, err
	}

	return state, true, nil
}

func (r *Repository) UpsertState(ctx context.Context, update FlightUpdate) error {
	observedAt := update.ObservedAt
	if observedAt.IsZero() {
		observedAt = time.Now().UTC()
	}

	query := `
		INSERT INTO flight_state (
			flight_id,
			status,
			departure_gate,
			departure_terminal,
			delay_minutes,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (flight_id) DO UPDATE SET
			status = EXCLUDED.status,
			departure_gate = EXCLUDED.departure_gate,
			departure_terminal = EXCLUDED.departure_terminal,
			delay_minutes = EXCLUDED.delay_minutes,
			updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		update.FlightID,
		update.Status,
		update.DepartureGate,
		update.DepartureTerminal,
		update.DelayMinutes,
		observedAt,
	)
	return err
}
