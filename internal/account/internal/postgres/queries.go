package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetByOwnerUUIDAndName(ctx context.Context, ownerUUID uuid.UUID, name string) (*account.Entity, error) {
	query := `
		SELECT a.id, a.uuid, a.name, a.user_id, a.created_at, a.updated_at
		FROM accounts a
		INNER JOIN users u ON a.user_id = u.id
		WHERE u.uuid = $1 AND a.name = $2
	`
	row := r.db.QueryRow(ctx, query, ownerUUID, name)

	accountEntity := &account.Entity{}

	err := row.Scan(
		&accountEntity.ID,
		&accountEntity.UUID,
		&accountEntity.Name,
		&accountEntity.UserID,
		&accountEntity.CreatedAt,
		&accountEntity.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, account.ErrAccountNotFound
		}

		return nil, fmt.Errorf(
			"repository->get by owner uuid and name: %w",
			err,
		)
	}

	return accountEntity, nil
}
