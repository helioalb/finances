package account

import (
	"context"
	"database/sql"

	"github.com/helioalb/finances/internal/user"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, account *Account) (*Account, error)
}

func Init(e *echo.Echo, db *sql.DB, user user.Service) Service {
	return nil
}
