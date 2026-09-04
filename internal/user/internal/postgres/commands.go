package postgres

import (
	"context"
	"fmt"

	"github.com/helioalb/finances/internal/user"
)

func (r *repository) Create(ctx context.Context, u *user.Entity) (*user.Entity, error) {
	if u == nil {
		return nil, fmt.Errorf("repository->user cannot be nil")
	}

	query := `INSERT INTO users
		(name, email)
	VALUES ($1, $2)
	RETURNING id, uuid, name, email, created_at,
		updated_at`
	row := r.db.QueryRow(ctx, query, u.Name, u.Email)

	createdUser := &user.Entity{}

	err := row.Scan(
		&createdUser.ID,
		&createdUser.UUID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"repository->create user: %w",
			err,
		)
	}

	return createdUser, nil
}
