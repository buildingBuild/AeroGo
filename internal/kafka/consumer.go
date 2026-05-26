package kafka

import (
	"context"
	"encoding/json"

	"aero-Go/internal/flights"
	kafkago "github.com/segmentio/kafka-go"
)

type FlightUpdateHandler func(context.Context, flights.FlightUpdate) error

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	return &Consumer{
		reader: kafkago.NewReader(kafkago.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

func (c *Consumer) Run(ctx context.Context, handler FlightUpdateHandler) error {
	for {
		message, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		var update flights.FlightUpdate
		if err := json.Unmarshal(message.Value, &update); err != nil {
			if commitErr := c.reader.CommitMessages(ctx, message); commitErr != nil {
				return commitErr
			}
			continue
		}

		if err := handler(ctx, update); err != nil {
			return err
		}

		if err := c.reader.CommitMessages(ctx, message); err != nil {
			return err
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
