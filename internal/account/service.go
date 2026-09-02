package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/user"
)

type repository interface {
	Create(ctx context.Context, account *Entity) (*Entity, error)
	GetByOwnerUUIDAndName(ctx context.Context, ownerUUID uuid.UUID, name string) (*Entity, error)
}

type userService interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*user.Entity, error)
}

type service struct {
	repo    repository
	userSvc userService
}

func newService(repo repository, userSvc userService) *service {
	return &service{
		repo:    repo,
		userSvc: userSvc,
	}
}

func (s *service) Create(ctx context.Context, input CreateInput) (*Entity, error) {
	_, err := s.repo.GetByOwnerUUIDAndName(ctx, input.OwnerUUID, input.Name)
	if err == nil {
		return nil, ErrAccountAlreadyExists
	}

	user, err := s.userSvc.GetByUUID(ctx, input.OwnerUUID)
	if err != nil {
		return nil, err
	}

	account := input.ToEntity(user.ID)

	createdAccount, err := s.repo.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	return createdAccount, nil
}
