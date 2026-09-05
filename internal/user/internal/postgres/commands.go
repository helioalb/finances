package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/helioalb/finances/internal/user"
	"github.com/jackc/pgx/v5/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
			return nil, fmt.Errorf("repository->create user: %w", user.ErrEmailInUse)
		}

		return nil, fmt.Errorf(
			"repository->create user: %w",
			err,
		)
	}

	return createdUser, nil
}
