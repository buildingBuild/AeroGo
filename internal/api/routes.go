package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", echo)
	mux.HandleFunc("/register-user", registerUser)
	mux.HandleFunc("/{user_id}/subscribe{flight_id}", echo)

	return mux
}
