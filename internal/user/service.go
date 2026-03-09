package user

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(
		ctx context.Context,
		user *User,
	) (*User, error)
	GetByEmail(
		ctx context.Context,
		email string,
	) (*User, error)
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

func (s *service) Create(ctx context.Context, user *User) (*User, error) {
	_, err := s.repo.GetByEmail(ctx, user.Email)
	if err == nil {
		return nil, ErrEmailInUse
	}

	user = user.New()

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service: create user: %w", err)
	}

	return createdUser, nil
}

func (s *service) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf(
			"service: get user by email: %w", err,
		)
	}

	return user, nil
}
