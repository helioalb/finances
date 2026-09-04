package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
)

func (r *repository) GetByOwnerUUIDAndName(ctx context.Context, ownerUUID uuid.UUID, name string) (*account.Entity, error) {
	query := `
		SELECT a.id, a.uuid, a.name, a.user_id, a.created_at, a.updated_at
		FROM accounts a
		INNER JOIN users u ON a.user_id = u.id
		WHERE u.uuid = $1 AND a.name = $2
	`
	row := r.db.QueryRow(ctx, query, ownerUUID, name)

	account := &account.Entity{}

	err := row.Scan(
		&account.ID,
		&account.UUID,
		&account.Name,
		&account.UserID,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"repository->get by owner uuid and name: %w",
			err,
		)
	}

	return account, nil
}
