package flights

import (
	"context"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ProcessUpdate(ctx context.Context, update FlightUpdate) ([]Change, error) {
	update = normalizeUpdate(update)

	if err := s.repo.EnsureFlight(ctx, update); err != nil {
		return nil, err
	}

	previous, found, err := s.repo.GetState(ctx, update.FlightID)
	if err != nil {
		return nil, err
	}

	changes := make([]Change, 0, 3)
	if found {
		changes = detectChanges(previous, update)
	}

	if err := s.repo.UpsertState(ctx, update); err != nil {
		return nil, err
	}

	return changes, nil
}

func normalizeUpdate(update FlightUpdate) FlightUpdate {
	update.FlightID = strings.ToUpper(strings.TrimSpace(update.FlightID))
	update.FlightNumber = strings.ToUpper(strings.TrimSpace(update.FlightNumber))
	update.Status = strings.ToLower(strings.TrimSpace(update.Status))
	update.DepartureGate = strings.ToUpper(strings.TrimSpace(update.DepartureGate))
	update.DepartureTerminal = strings.ToUpper(strings.TrimSpace(update.DepartureTerminal))

	if update.FlightNumber == "" {
		update.FlightNumber = update.FlightID
	}
	if update.EventID == "" {
		update.EventID = update.FlightID + "-" + time.Now().UTC().Format("20060102150405.000000000")
	}
	if update.ObservedAt.IsZero() {
		update.ObservedAt = time.Now().UTC()
	}

	return update
}

func detectChanges(previous FlightState, update FlightUpdate) []Change {
	var changes []Change

	if previous.DepartureGate != update.DepartureGate {
		changes = append(changes, Change{
			Type:     "gate_change",
			OldValue: previous.DepartureGate,
			NewValue: update.DepartureGate,
		})
	}
	if previous.DelayMinutes != update.DelayMinutes {
		changes = append(changes, Change{
			Type:     "delay_change",
			OldValue: intToString(previous.DelayMinutes),
			NewValue: intToString(update.DelayMinutes),
		})
	}
	if previous.Status != update.Status {
		changes = append(changes, Change{
			Type:     "status_change",
			OldValue: previous.Status,
			NewValue: update.Status,
		})
	}

	return changes
}

func intToString(value int) string {
	return strconv.Itoa(value)
}
