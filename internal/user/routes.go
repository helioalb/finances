package user

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo, h *handler) {
	e.POST("/users", h.Create)
}
