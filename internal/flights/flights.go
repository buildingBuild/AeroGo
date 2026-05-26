package flights

import "time"

type FlightUpdate struct {
	EventID           string    `json:"event_id"`
	FlightID          string    `json:"flight_id"`
	FlightNumber      string    `json:"flight_number"`
	Status            string    `json:"status"`
	DepartureGate     string    `json:"departure_gate"`
	DepartureTerminal string    `json:"departure_terminal"`
	DelayMinutes      int       `json:"delay_minutes"`
	ObservedAt        time.Time `json:"observed_at"`
}

type FlightState struct {
	FlightID          string
	Status            string
	DepartureGate     string
	DepartureTerminal string
	DelayMinutes      int
	UpdatedAt         time.Time
}

type Change struct {
	Type     string
	OldValue string
	NewValue string
}
