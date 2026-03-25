package account

import "github.com/labstack/echo"

func registerRoutes(e *echo.Echo, account *handler) {
	e.POST("/accounts", account.Create)
}
