package transaction

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
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
	var input CreateInput

	if err := c.Bind(&input); err != nil {
		return h.badRequestResponse(c, err)
	}

	if err := input.Validate(); err != nil {
		return h.unprocessableEntityResponse(c, err)
	}

	ctx := c.Request().Context()

	accountUUID, err := uuid.Parse(input.AccountUUID)
	if err != nil {
		return h.badRequestResponse(c, err)
	}

	err = h.svc.Expense(ctx, accountUUID, input.Amount, input.Description)
	if err != nil {
		if errors.Is(err, errAccountNotFound) {
			return h.accountNotFoundResponse(c)
		}

		return h.internalServerErrorResponse(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *handler) accountNotFoundResponse(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "account not found",
	})
}

func (h *handler) badRequestResponse(c echo.Context, err error) error {
	return c.JSON(400, map[string]string{
		"error": err.Error(),
	})
}

func (h *handler) unprocessableEntityResponse(c echo.Context, err error) error {
	return c.JSON(422, map[string]string{
		"error": err.Error(),
	})
}

func (h *handler) internalServerErrorResponse(c echo.Context, err error) error {
	h.log.Error(
		"[transaction][expense]",
		"[error=", err.Error(), "]",
	)

	return c.JSON(500, map[string]string{
		"error": "internal server error",
	})
}

func (h *handler) Income(c echo.Context) error {
	return nil
}

func (h *handler) Transfer(c echo.Context) error {
	return nil
}
