package main

import (
	"aero-Go/internal/api"
	"aero-Go/internal/db"
	"aero-Go/internal/subscriptions"
	"aero-Go/internal/users"
	"fmt"
	"log"
	"net/http"
)

func main() {

	database := db.StartDbConnection()
	defer database.Close()

	userRepo := users.NewRepository(database)
	userService := users.NewService(userRepo)

	subscriptionRepo := subscriptions.NewRepository(database)
	subscriptionService := subscriptions.NewService(subscriptionRepo)

	router := api.NewRouter(userService, subscriptionService)

	fmt.Println("Base landing")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}

}
