package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgRepository struct {
	db *pgxpool.Pool
}

func newPgRepository(db *pgxpool.Pool) *pgRepository {
	return &pgRepository{db: db}
}

func (r *pgRepository) Create(ctx context.Context, accountUUID uuid.UUID, transaction *Entity) error {

	query := `INSERT INTO transactions (account_id, amount, type, description)
	 SELECT a.id, $1, $2, $3 FROM accounts a WHERE a.uuid = $4
	 RETURNING uuid, account_id, amount, type, created_at;
	 `

	row := r.db.QueryRow(
		ctx,
		query,
		transaction.Amount,
		transaction.Type,
		transaction.Description,
		accountUUID,
	)

	createdTransaction := &Entity{}

	err := row.Scan(
		&createdTransaction.UUID,
		&createdTransaction.AccountID,
		&createdTransaction.Amount,
		&createdTransaction.Type,
		&createdTransaction.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errAccountNotFound
		}

		return fmt.Errorf(
			"repository->create transaction: %w",
			err,
		)
	}

	return nil
}
