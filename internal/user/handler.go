package user

import (
	"errors"
	"net/http"

	"github.com/helioalb/finances/internal/platform/httpx"
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

func (h *handler) Create(c echo.Context) error {
	var input CreateInput
	requestID := httpx.RequestID(c)

	if err := c.Bind(&input); err != nil {
		return h.badRequestResponse(c, err)
	}

	if err := input.Validate(); err != nil {
		return h.unprocessableEntityResponse(c, err)
	}

	ctx := c.Request().Context()

	user, err := h.svc.Create(ctx, input)
	if err != nil {
		if errors.Is(err, ErrEmailInUse) {
			return h.emailAlreadyInUseResponse(c)
		}

		return h.internalServerErrorResponse(c, err)
	}

	h.log.Info(
		"[user][create]",
		"[http_status=", http.StatusCreated, "]",
		"[user_uuid=", user.UUID.String(), "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusCreated, map[string]string{
		"uuid": user.UUID.String(),
	})
}

func (h *handler) badRequestResponse(c echo.Context, err error) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[user][create]",
		"[http_status=", http.StatusBadRequest, "]",
		"[error=", err.Error(), "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": err.Error()},
	)
}

func (h *handler) unprocessableEntityResponse(c echo.Context, err error) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[user][create]",
		"[http_status=", http.StatusUnprocessableEntity, "]",
		"[error=", err.Error(), "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusUnprocessableEntity, map[string]string{
		"error": err.Error()},
	)
}

func (h *handler) emailAlreadyInUseResponse(c echo.Context) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[user][create][email_already_in_use]",
		"[http_status=", http.StatusConflict, "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusConflict, map[string]string{
		"error": "email already in use"},
	)
}

func (h *handler) internalServerErrorResponse(c echo.Context, err error) error {
	requestID := httpx.RequestID(c)

	h.log.Error(
		"[user][create]",
		"[http_status=", http.StatusInternalServerError, "]",
		"[err=", err.Error(), "]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "internal server error"},
	)
}
