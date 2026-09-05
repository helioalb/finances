package factory

import (
	"github.com/helioalb/finances/internal/account"
	"github.com/helioalb/finances/internal/account/internal/postgres"
	"github.com/helioalb/finances/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BuildService(db *pgxpool.Pool, userSvc user.Service) account.Service {
	repo := postgres.NewRepository(db)
	return account.NewService(repo, userSvc)
}
