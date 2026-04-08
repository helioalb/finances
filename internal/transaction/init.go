package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

type Service interface {
	Expense(ctx context.Context, accountUUID uuid.UUID, amount int, description *string) error
	Income(ctx context.Context, accountUUID uuid.UUID, amount int, description *string) error
	Transfer(ctx context.Context, fromAccountUUID uuid.UUID, toAccountUUID uuid.UUID, amount int) error
}

func Init(e *echo.Echo, db *pgxpool.Pool) Service {
	repo := newPgRepository(db)
	svc := newService(repo)
	handler := newHandler(svc, e.Logger)

	registerRoutes(e, handler)

	return svc
}
