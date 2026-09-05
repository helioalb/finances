package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/helioalb/finances/internal/account"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) Create(ctx context.Context, a *account.Entity) (*account.Entity, error) {
	if a == nil {
		return nil, fmt.Errorf("repository->account cannot be nil")
	}

	query := `INSERT INTO accounts (name, user_id, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, uuid, name, user_id, created_at, updated_at`

	row := r.db.QueryRow(ctx, query, a.Name, a.UserID)
	createdAccount := &account.Entity{}

	err := row.Scan(
		&createdAccount.ID,
		&createdAccount.UUID,
		&createdAccount.Name,
		&createdAccount.UserID,
		&createdAccount.CreatedAt,
		&createdAccount.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "uq_accounts_user_name" {
			return nil, fmt.Errorf("repository->create account: %w", account.ErrAccountAlreadyExists)
		}

		return nil, fmt.Errorf(
			"repository->create account: %w",
			err,
		)
	}

	return createdAccount, nil
}
