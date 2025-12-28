package user

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	if repo == nil {
		panic("repo cannot be nil")
	}
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name string) (*User, error) {
	user := Create(name)

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service: create user: %w", err)
	}

	return createdUser, nil
}
