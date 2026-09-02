package account

import (
	"context"

	"github.com/helioalb/finances/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
}

func Init(db *pgxpool.Pool, userSvc user.Service, log echo.Logger) Service {
	repo := newPgRepository(db)
	return newService(repo, userSvc)
}
