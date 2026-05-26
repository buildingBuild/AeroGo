package notifications

import (
	"context"
	"fmt"

	"aero-Go/internal/alerts"
)

type ConsoleSender struct{}

func NewConsoleSender() ConsoleSender {
	return ConsoleSender{}
}

func (s ConsoleSender) SendAlert(ctx context.Context, alert alerts.Alert) error {
	_ = ctx
	fmt.Printf("SMS alert for user %s: %s\n", alert.UserID, alert.Message)
	return nil
}
