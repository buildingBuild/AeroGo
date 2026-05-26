package kafka

import (
	"context"
	"encoding/json"
	"time"

	"aero-Go/internal/flights"
	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkago.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafkago.Writer{
			Addr:         kafkago.TCP(brokers...),
			Topic:        topic,
			RequiredAcks: kafkago.RequireOne,
			Balancer:     &kafkago.Hash{},
			BatchTimeout: 10 * time.Millisecond,
		},
	}
}

func (p *Producer) PublishFlightUpdate(ctx context.Context, update flights.FlightUpdate) error {
	payload, err := json.Marshal(update)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafkago.Message{
		Key:   []byte(update.FlightID),
		Value: payload,
		Time:  update.ObservedAt,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
