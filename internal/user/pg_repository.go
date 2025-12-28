package user

import (
	"context"
	"database/sql"
	"fmt"
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
		return nil, fmt.Errorf("repository: user cannot be nil")
	}

	query := `INSERT INTO users (uuid, name, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, uuid, name, created_at, updated_at`
	row := r.db.QueryRowContext(ctx, query, user.UUID, user.Name)

	createdUser := &User{}

	err := row.Scan(&createdUser.ID, &createdUser.UUID, &createdUser.Name, &createdUser.CreatedAt, &createdUser.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("repository: create user: %w, query: %s", err, query)
	}

	return createdUser, nil
}
