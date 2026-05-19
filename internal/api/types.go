package api

type UserRegistration struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
