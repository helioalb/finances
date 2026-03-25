package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error)
}

func Init(e *echo.Echo, db *sql.DB, log echo.Logger) Service {
	repo := newPgRepository(db)
	svc := newService(repo)
	handler := newHandler(svc, log)

	registerRoutes(e, handler)

	return svc
}
