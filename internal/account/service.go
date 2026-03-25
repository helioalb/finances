package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/user"
)

type repository interface {
	Create(ctx context.Context, account *Account) (*Account, error)
	GetByOwnerUUIDAndName(ctx context.Context, ownerUUID uuid.UUID, name string) (*Account, error)
}

type service struct {
	repo    repository
	userSvc user.Service
}

func newService(repo repository, userSvc user.Service) *service {
	return &service{
		repo:    repo,
		userSvc: userSvc,
	}
}

func (s *service) Create(ctx context.Context, input CreateInput) (*Account, error) {
	_, err := s.repo.GetByOwnerUUIDAndName(ctx, input.OwnerUUID, input.Name)
	if err == nil {
		return nil, errAccountAlreadyExists
	}

	user, err := s.userSvc.GetByUUID(ctx, input.OwnerUUID)
	if err != nil {
		return nil, err
	}

	account := &Account{
		Name:   input.Name,
		UserID: user.ID,
	}

	createdAccount, err := s.repo.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	return createdAccount, nil
}
