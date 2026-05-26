package api

import (
	"aero-Go/internal/subscriptions"
	"aero-Go/internal/users"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	userService         *users.Service
	subscriptionService *subscriptions.Service
}

func NewHandler(userService *users.Service, subscriptionService *subscriptions.Service) *Handler {
	return &Handler{
		userService:         userService,
		subscriptionService: subscriptionService,
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good")
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body UserRegistration

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), users.CreateUserInput{
		Name:        body.Name,
		PhoneNumber: body.Number,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) subscribeToFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body FlightSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	subscription, err := h.subscriptionService.Subscribe(r.Context(), subscriptions.CreateSubscriptionInput{
		UserID:   body.UserID,
		FlightID: body.FlightID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subscription)
}
