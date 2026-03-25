package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type repository interface {
	Create(ctx context.Context, user *Entity) (*Entity, error)
	GetByEmail(ctx context.Context, email string) (*Entity, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error)
}

type service struct {
	repo repository
}

func newService(repo repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, input CreateInput) (*Entity, error) {
	_, err := s.repo.GetByEmail(ctx, input.Email)
	if err == nil {
		return nil, errEmailInUse
	}

	user := input.ToEntity()

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service->%w", err)
	}

	return createdUser, nil
}

func (s *service) GetByEmail(ctx context.Context, email string) (*Entity, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf(
			"service->%w", err,
		)
	}

	return user, nil
}

func (s *service) GetByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error) {
	user, err := s.repo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(
			"service->%w", err,
		)
	}

	return user, nil
}
