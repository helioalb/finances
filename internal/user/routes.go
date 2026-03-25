package user

import "github.com/labstack/echo"

func registerRoutes(e *echo.Echo, user *handler) {
	e.POST("/users", user.Create)
}
