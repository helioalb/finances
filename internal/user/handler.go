package user

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	service Service
	logger  *slog.Logger
}

func newHandler(service Service, logger *slog.Logger) *handler {
	if service == nil {
		panic("service cannot be nil")
	}

	if logger == nil {
		panic("logger cannot be nil")
	}

	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Create(c echo.Context) error {
	var req createInput

	if err := c.Bind(&req); err != nil {
		return badRequestResponse(c, h, err)
	}

	if err := req.Validate(); err != nil {
		return unprocessableEntityResponse(c, h, err)
	}

	ctx := c.Request().Context()

	user, err := h.service.Create(ctx, req.ToEntity())
	if err != nil {
		if errors.Is(err, ErrEmailInUse) {
			return emailAlreadyInUseResponse(c, h, req)
		}

		return internalServerErrorResponse(c, h, req, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"uuid": user.UUID.String(),
	})
}

func badRequestResponse(c echo.Context, h *handler, err error) error {
	h.logger.Warn(
		fmt.Sprintf("[user][create][bad_request][error=%v]", err),
	)

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "bad request"},
	)
}

func unprocessableEntityResponse(c echo.Context, h *handler, err error) error {
	h.logger.Warn(
		fmt.Sprintf("[user][create][unprocessable_entity][error=%v]", err),
	)

	return c.JSON(http.StatusUnprocessableEntity, map[string]string{
		"error": "unprocessable entity"},
	)
}

func emailAlreadyInUseResponse(c echo.Context, h *handler, req createInput) error {
	h.logger.Warn(
		fmt.Sprintf("[user][create][email_already_in_use][email=%s]", req.Email),
	)

	return c.JSON(http.StatusConflict, map[string]string{
		"error": "email already in use"},
	)
}

func internalServerErrorResponse(c echo.Context, h *handler, req createInput, err error) error {
	h.logger.Error(
		fmt.Sprintf("[user][create][error=%v][email=%s]", err, req.Email),
	)

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "internal server error"},
	)
}
