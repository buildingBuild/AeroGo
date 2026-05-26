package alerts

import (
	"context"
	"fmt"
	"strings"

	"aero-Go/internal/flights"
)

type Notifier interface {
	SendAlert(ctx context.Context, alert Alert) error
}

type Service struct {
	repo     *Repository
	notifier Notifier
}

func NewService(repo *Repository, notifier Notifier) *Service {
	return &Service{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *Service) HandleChanges(ctx context.Context, update flights.FlightUpdate, changes []flights.Change) error {
	for _, change := range changes {
		eventKey := buildEventKey(update.FlightID, change)
		message := buildMessage(update, change)

		alerts, err := s.repo.CreateForSubscribers(ctx, update.FlightID, change.Type, eventKey, message)
		if err != nil {
			return err
		}

		for _, alert := range alerts {
			if err := s.notifier.SendAlert(ctx, alert); err != nil {
				return err
			}
			if err := s.repo.MarkSent(ctx, alert.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

func buildEventKey(flightID string, change flights.Change) string {
	return strings.Join([]string{flightID, change.Type, change.NewValue}, ":")
}

func buildMessage(update flights.FlightUpdate, change flights.Change) string {
	switch change.Type {
	case "gate_change":
		return fmt.Sprintf("%s gate changed from %s to %s", update.FlightNumber, emptyValue(change.OldValue), change.NewValue)
	case "delay_change":
		return fmt.Sprintf("%s delay changed from %s to %s minutes", update.FlightNumber, change.OldValue, change.NewValue)
	case "status_change":
		return fmt.Sprintf("%s status changed from %s to %s", update.FlightNumber, emptyValue(change.OldValue), change.NewValue)
	default:
		return fmt.Sprintf("%s changed: %s -> %s", update.FlightNumber, emptyValue(change.OldValue), change.NewValue)
	}
}

func emptyValue(value string) string {
	if value == "" {
		return "unknown"
	}
	return value
}
