package transaction

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func newPgRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, transaction *Transaction) error {
	return nil
}
