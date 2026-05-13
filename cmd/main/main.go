package main

import (
	"aero-Go/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()

	fmt.Println("Base landing")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
