package account

import (
	"context"
	"database/sql"

	"github.com/helioalb/finances/internal/user"
	"github.com/labstack/echo"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
}

func Init(e *echo.Echo, db *sql.DB, userSvc user.Service, log echo.Logger) Service {
	repo := newPgRepository(db)
	service := newService(repo, userSvc)
	handler := newHandler(service, log)

	registerRoutes(e, handler)

	return service
}
