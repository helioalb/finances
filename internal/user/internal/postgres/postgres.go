package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &repository{db: db}
}
