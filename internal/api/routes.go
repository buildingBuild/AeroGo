package api

import (
	"aero-Go/internal/subscriptions"
	"aero-Go/internal/users"
	"net/http"
)

func NewRouter(userService *users.Service, subscriptionService *subscriptions.Service) http.Handler {
	handler := NewHandler(userService, subscriptionService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", echo)
	mux.HandleFunc("/register-user", handler.registerUser)
	mux.HandleFunc("/subscriptions", handler.subscribeToFlight)

	return mux
}
