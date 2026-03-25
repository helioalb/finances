package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*User, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*User, error)
}

func Init(e *echo.Echo, db *sql.DB, log echo.Logger) Service {
	repo := newPgRepository(db)
	svc := newService(repo)
	handler := newHandler(svc, log)

	RegisterRoutes(e, handler)

	return svc
}
