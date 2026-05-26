package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	aerokafka "aero-Go/internal/kafka"
	"aero-Go/internal/poller"
)

func main() {
	brokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	topic := getEnv("KAFKA_TOPIC", "flight-updates")

	producer := aerokafka.NewProducer(brokers, topic)
	defer producer.Close()

	ctx := context.Background()
	for _, update := range poller.MockUpdates() {
		if err := producer.PublishFlightUpdate(ctx, update); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("published %s for %s\n", update.EventID, update.FlightID)
		time.Sleep(750 * time.Millisecond)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
