package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/transaction"
	"github.com/jackc/pgx/v5"
)

func (r *repository) Create(ctx context.Context, accountUUID uuid.UUID, t *transaction.Entity) error {

	query := `INSERT INTO transactions (account_id, amount, type, description)
	 SELECT a.id, $1, $2, $3 FROM accounts a WHERE a.uuid = $4
	 RETURNING uuid, account_id, amount, type, created_at;
	 `

	row := r.db.QueryRow(
		ctx,
		query,
		t.Amount,
		t.Type,
		t.Description,
		accountUUID,
	)

	createdTransaction := &transaction.Entity{}

	err := row.Scan(
		&createdTransaction.UUID,
		&createdTransaction.AccountID,
		&createdTransaction.Amount,
		&createdTransaction.Type,
		&createdTransaction.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return transaction.ErrAccountNotFound
		}

		return fmt.Errorf(
			"repository->create transaction: %w",
			err,
		)
	}

	return nil
}
