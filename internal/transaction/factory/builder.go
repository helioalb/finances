package factory

import (
	"github.com/helioalb/finances/internal/transaction"
	"github.com/helioalb/finances/internal/transaction/internal/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BuildService(db *pgxpool.Pool) transaction.Service {
	repo := postgres.NewRepository(db)
	return transaction.NewService(repo)
}
