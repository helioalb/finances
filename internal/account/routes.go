package account

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo, account *handler) {
	e.POST("/accounts", account.Create)
}
