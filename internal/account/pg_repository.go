package account

import (
	"database/sql"
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

// func Create(ctx context.Context, account *Account) (*Account, error) {
// 	if account == nil {
// 		return nil, fmt.Errorf("repository->account cannot be nil")
// 	}
// }
