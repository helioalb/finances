package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/user"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetByEmail(ctx context.Context, email string) (*user.Entity, error) {
	query := `
		SELECT id, uuid, name, email, created_at, updated_at
		FROM users WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, email)

	u := &user.Entity{}

	err := row.Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf(
			"repository->get user by email: %w",
			err,
		)
	}

	return u, nil
}

func (r *repository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*user.Entity, error) {
	query := `
		SELECT id, uuid, name, email, created_at, updated_at
		FROM users WHERE uuid = $1
	`
	row := r.db.QueryRow(ctx, query, uuid)

	u := &user.Entity{}

	err := row.Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf(
			"repository->get user by uuid: %w",
			err,
		)
	}

	return u, nil
}
