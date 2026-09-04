package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	if db == nil {
		panic("db cannot be nil")
	}

	return &repository{db: db}
}

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
		return nil, fmt.Errorf(
			"repository->create account: %w",
			err,
		)
	}

	return createdAccount, nil
}

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
