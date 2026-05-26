package api

type UserRegistration struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type FlightSubscriptionRequest struct {
	UserID   string `json:"user_id"`
	FlightID string `json:"flight_id"`
}
