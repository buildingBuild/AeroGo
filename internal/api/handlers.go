package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var body UserRegistration

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

}
