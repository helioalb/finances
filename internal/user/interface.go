package user

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error)
}
