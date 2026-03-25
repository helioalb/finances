package user

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo, user *handler) {
	e.POST("/users", user.Create)
}
