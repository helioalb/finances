package transaction

import (
	"context"

	"github.com/google/uuid"
)

type repository interface {
	Create(ctx context.Context, accountUUID uuid.UUID, transaction *Entity) error
}

type service struct {
	repo repository
}

func NewService(repo repository) *service {
	return &service{repo: repo}
}

func (s *service) Expense(ctx context.Context, accountUUID uuid.UUID, amount int, description *string) error {
	transaction := &Entity{
		Amount:      amount * -1,
		Type:        "EXPENSE",
		Description: description,
	}

	return s.repo.Create(ctx, accountUUID, transaction)
}

func (s *service) Income(ctx context.Context, accountUUID uuid.UUID, amount int, description *string) error {
	transaction := &Entity{
		Amount:      amount,
		Type:        "INCOME",
		Description: description,
	}

	return s.repo.Create(ctx, accountUUID, transaction)
}

func (s *service) Transfer(ctx context.Context, fromAccountUUID uuid.UUID, toAccountUUID uuid.UUID, amount int) error {
	return nil
}
