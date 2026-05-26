package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"aero-Go/internal/alerts"
	"aero-Go/internal/db"
	"aero-Go/internal/flights"
	aerokafka "aero-Go/internal/kafka"
	"aero-Go/internal/notifications"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	database := db.StartDbConnection()
	defer database.Close()

	flightRepo := flights.NewRepository(database)
	flightService := flights.NewService(flightRepo)

	alertRepo := alerts.NewRepository(database)
	notifier := notifications.NewRetrySender(notifications.NewConsoleSender(), 3, time.Second)
	alertService := alerts.NewService(alertRepo, notifier)

	brokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	topic := getEnv("KAFKA_TOPIC", "flight-updates")
	groupID := getEnv("KAFKA_GROUP_ID", "aerogo-alerts")

	consumer := aerokafka.NewConsumer(brokers, topic, groupID)
	defer consumer.Close()

	fmt.Printf("consuming %s from %s\n", topic, strings.Join(brokers, ","))
	err := consumer.Run(ctx, func(ctx context.Context, update flights.FlightUpdate) error {
		changes, err := flightService.ProcessUpdate(ctx, update)
		if err != nil {
			return err
		}

		if len(changes) == 0 {
			fmt.Printf("stored baseline for %s\n", update.FlightID)
			return nil
		}

		if err := alertService.HandleChanges(ctx, update, changes); err != nil {
			return err
		}

		fmt.Printf("processed %d change(s) for %s\n", len(changes), update.FlightID)
		return nil
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
