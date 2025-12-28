package user

import (
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
	type requestBody struct {
		Name string `json:"name"`
	}

	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.service.Create(c.Request().Context(), body.Name)
	if err != nil {
		h.logger.Error("failed to create user", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"uuid": user.UUID.String()})
}

func RegisterRoutes(e *echo.Echo, handler *handler) {
	e.POST("/users", handler.Create)
}
