package transaction

import (
	"github.com/labstack/echo"
)

type handler struct {
	svc Service
	log echo.Logger
}

func newHandler(service Service, log echo.Logger) *handler {
	return &handler{
		svc: service,
		log: log,
	}
}

func (h *handler) Expense(c echo.Context) error {
	return nil
}

func (h *handler) Income(c echo.Context) error {
	return nil
}

func (h *handler) Transfer(c echo.Context) error {
	return nil
}
