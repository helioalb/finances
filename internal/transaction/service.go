package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
)

type repository interface {
	Create(ctx context.Context, transaction *Entity) error
}

type service struct {
	repo       repository
	accountSvc account.Service
}

func newService(repo repository, accountSvc account.Service) *service {
	return &service{repo: repo, accountSvc: accountSvc}
}

func (s *service) Expense(ctx context.Context, accountUUID uuid.UUID, amount int) error {
	return nil
}

func (s *service) Income(ctx context.Context, accountUUID uuid.UUID, amount int) error {
	return nil
}

func (s *service) Transfer(ctx context.Context, fromAccountUUID uuid.UUID, toAccountUUID uuid.UUID, amount int) error {
	return nil
}
