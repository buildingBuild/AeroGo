package api

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", echo)

	return mux
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good")
}
