package poller

import (
	"time"

	"aero-Go/internal/flights"
)

func MockUpdates() []flights.FlightUpdate {
	now := time.Now().UTC()

	return []flights.FlightUpdate{
		{
			EventID:           "AA123-001",
			FlightID:          "AA123",
			FlightNumber:      "AA123",
			Status:            "scheduled",
			DepartureGate:     "A4",
			DepartureTerminal: "1",
			DelayMinutes:      0,
			ObservedAt:        now,
		},
		{
			EventID:           "AA123-002",
			FlightID:          "AA123",
			FlightNumber:      "AA123",
			Status:            "scheduled",
			DepartureGate:     "A7",
			DepartureTerminal: "1",
			DelayMinutes:      0,
			ObservedAt:        now.Add(1 * time.Minute),
		},
		{
			EventID:           "AA123-003",
			FlightID:          "AA123",
			FlightNumber:      "AA123",
			Status:            "delayed",
			DepartureGate:     "A7",
			DepartureTerminal: "1",
			DelayMinutes:      35,
			ObservedAt:        now.Add(2 * time.Minute),
		},
		{
			EventID:           "AA123-004",
			FlightID:          "AA123",
			FlightNumber:      "AA123",
			Status:            "departed",
			DepartureGate:     "A7",
			DepartureTerminal: "1",
			DelayMinutes:      35,
			ObservedAt:        now.Add(3 * time.Minute),
		},
	}
}
