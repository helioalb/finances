package transaction

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
	"github.com/labstack/echo"
)

type Service interface {
	Expense(ctx context.Context, accountUUID uuid.UUID, amount int) error
	Income(ctx context.Context, accountUUID uuid.UUID, amount int) error
	Transfer(ctx context.Context, fromAccountUUID uuid.UUID, toAccountUUID uuid.UUID, amount int) error
}

func Init(e *echo.Echo, db *sql.DB, account account.Service) Service {
	_ = account

	repo := newPgRepository(db)
	svc := newService(repo)
	handler := newHandler(svc, e.Logger)

	RegisterRoutes(e, handler)

	return svc
}
