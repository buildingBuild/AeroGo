package subscriptions

import "time"

type Subscription struct {
	UserID    string    `json:"user_id"`
	FlightID  string    `json:"flight_id"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSubscriptionInput struct {
	UserID   string
	FlightID string
}
