package transaction

import (
	"context"

	"github.com/google/uuid"
)

type repository interface {
	Create(ctx context.Context, transaction *Transaction) error
}

type service struct {
	repo repository
}

func newService(repo repository) *service {
	return &service{repo: repo}
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
