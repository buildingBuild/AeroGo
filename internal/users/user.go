package users

import "time"

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"number"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserInput struct {
	Name        string
	PhoneNumber string
}
