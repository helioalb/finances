package user

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) *service {
	if repo == nil {
		panic("repo cannot be nil")
	}
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, name string) (*User, error) {
	user := Create(name)

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service: create user: %w", err)
	}

	return createdUser, nil
}
