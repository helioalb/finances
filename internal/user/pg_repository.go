package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type repository struct {
	db *sql.DB
}

func newPgRepository(db *sql.DB) *repository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) (*User, error) {
	if user == nil {
		return nil, fmt.Errorf("repository->user cannot be nil")
	}

	query := `INSERT INTO users
		(name, email)
	VALUES ($1, $2)
	RETURNING id, uuid, name, email, created_at,
		updated_at`
	row := r.db.QueryRowContext(ctx, query, user.Name, user.Email)

	createdUser := &User{}

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

func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, uuid, name, email, created_at, updated_at
		FROM users WHERE email = $1
	`
	row := r.db.QueryRowContext(ctx, query, email)

	user := &User{}

	err := row.Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf(
			"repository->get user by email: %w",
			err,
		)
	}

	return user, nil
}

func (r *repository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*User, error) {
	query := `
		SELECT id, uuid, name, email, created_at, updated_at
		FROM users WHERE uuid = $1
	`
	row := r.db.QueryRowContext(ctx, query, uuid)

	user := &User{}

	err := row.Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf(
			"repository->get user by uuid: %w",
			err,
		)
	}

	return user, nil
}
