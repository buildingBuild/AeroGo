package users

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, input CreateUserInput) (User, error) {
	input.Name = strings.TrimSpace(input.Name)
	input.PhoneNumber = strings.TrimSpace(input.PhoneNumber)

	if input.Name == "" {
		return User{}, errors.New("name is required")
	}
	if input.PhoneNumber == "" {
		return User{}, errors.New("number is required")
	}

	return s.repo.Create(ctx, input)
}
