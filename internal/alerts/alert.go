package alerts

import "time"

type Alert struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	FlightID  string    `json:"flight_id"`
	Type      string    `json:"alert_type"`
	Message   string    `json:"message"`
	EventKey  string    `json:"event_key"`
	CreatedAt time.Time `json:"created_at"`
}
