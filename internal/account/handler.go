package account

import (
	"errors"
	"net/http"

	"github.com/helioalb/finances/internal/platform/httpx"
	"github.com/labstack/echo"
)

type handler struct {
	service Service
	log     echo.Logger
}

func newHandler(service Service, log echo.Logger) *handler {
	return &handler{
		service: service,
		log:     log,
	}
}

func (h *handler) Create(c echo.Context) error {
	var input CreateInput

	if err := c.Bind(&input); err != nil {
		return h.badRequestResponse(c, err)
	}

	if err := input.Validate(); err != nil {
		return h.unprocessableEntityResponse(c, err)
	}

	ctx := c.Request().Context()

	account, err := h.service.Create(ctx, input)
	if err != nil {
		if errors.Is(err, errAccountAlreadyExists) {
			return h.accountAlreadyExistsResponse(c)
		}

		return h.internalServerErrorResponse(c, err)
	}

	h.log.Info(
		"[account][create]",
		"[http_status=", http.StatusCreated, "]",
		"[account_uuid=", account.UUID.String(), "]",
	)

	return c.JSON(http.StatusCreated, map[string]string{
		"uuid": account.UUID.String(),
	})
}

func (h *handler) badRequestResponse(c echo.Context, err error) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[account][create]",
		"[http_status=", http.StatusBadRequest, "]",
		"[error=", err.Error(), "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": err.Error(),
	})
}

func (h *handler) unprocessableEntityResponse(c echo.Context, err error) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[account][create]",
		"[http_status=", http.StatusUnprocessableEntity, "]",
		"[error=", err.Error(), "]",
		"[request_id=", requestID, "]",
	)
	return c.JSON(http.StatusUnprocessableEntity, map[string]string{
		"error": err.Error(),
	})
}

func (h *handler) accountAlreadyExistsResponse(c echo.Context) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[account][create]",
		"[http_status=", http.StatusConflict, "]",
		"[error=account already exists]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusConflict, map[string]string{
		"error": "account already exists",
	})
}

func (h *handler) internalServerErrorResponse(c echo.Context, err error) error {
	h.log.Error(
		"[account][create]",
		"[http_status=", http.StatusInternalServerError, "]",
		"[error=", err.Error(), "]",
	)

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "internal server error",
	})
}
