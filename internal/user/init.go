package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error)
}

func Init(db *pgxpool.Pool, log echo.Logger) Service {
	repo := newPgRepository(db)
	return newService(repo)
}
