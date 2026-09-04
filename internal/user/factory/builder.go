package factory

import (
	"github.com/helioalb/finances/internal/user"
	"github.com/helioalb/finances/internal/user/internal/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

func BuildService(db *pgxpool.Pool, log echo.Logger) user.Service {
	repo := postgres.NewRepository(db)
	return user.NewService(repo)
}
