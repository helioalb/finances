package factory

import (
	"github.com/helioalb/finances/internal/account"
	"github.com/helioalb/finances/internal/account/internal/postgres"
	"github.com/helioalb/finances/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

func BuildService(db *pgxpool.Pool, userSvc user.Service, log echo.Logger) account.Service {
	repo := postgres.NewRepository(db)
	return account.NewService(repo, userSvc)
}
