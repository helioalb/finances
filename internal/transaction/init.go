package transaction

import (
	"context"
	"database/sql"

	"github.com/helioalb/finances/internal/account"
	"github.com/labstack/echo"
)

type Service interface {
	Expense(ctx context.Context) error
}

func Init(e *echo.Echo, db *sql.DB, account account.Service) Service {
	return nil
}
