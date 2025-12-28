package user

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, name string) (*User, error)
}

func Init(e *echo.Echo, db *sql.DB, logger *slog.Logger) Service {
	repo := newPgRepository(db)
	svc := newService(repo)
	handler := newHandler(svc, logger)

	RegisterRoutes(e, handler)

	return svc
}
