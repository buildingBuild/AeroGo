package notifications

import (
	"context"
	"time"

	"aero-Go/internal/alerts"
)

type RetrySender struct {
	sender   alerts.Notifier
	attempts int
	delay    time.Duration
}

func NewRetrySender(sender alerts.Notifier, attempts int, delay time.Duration) RetrySender {
	return RetrySender{
		sender:   sender,
		attempts: attempts,
		delay:    delay,
	}
}

func (s RetrySender) SendAlert(ctx context.Context, alert alerts.Alert) error {
	var lastErr error

	for attempt := 1; attempt <= s.attempts; attempt++ {
		if err := s.sender.SendAlert(ctx, alert); err != nil {
			lastErr = err
			time.Sleep(s.delay)
			continue
		}
		return nil
	}

	return lastErr
}
