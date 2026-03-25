package transaction

import (
	"context"
	"database/sql"
)

type pgRepository struct {
	db *sql.DB
}

func newPgRepository(db *sql.DB) *pgRepository {
	return &pgRepository{db: db}
}

func (r *pgRepository) Create(ctx context.Context, transaction *Entity) error {
	return nil
}
