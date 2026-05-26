package subscriptions

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Subscribe(ctx context.Context, input CreateSubscriptionInput) (Subscription, error) {
	input.UserID = strings.TrimSpace(input.UserID)
	input.FlightID = strings.ToUpper(strings.TrimSpace(input.FlightID))

	if input.UserID == "" {
		return Subscription{}, errors.New("user_id is required")
	}
	if input.FlightID == "" {
		return Subscription{}, errors.New("flight_id is required")
	}

	return s.repo.Create(ctx, input)
}
