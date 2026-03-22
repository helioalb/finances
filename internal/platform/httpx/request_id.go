package httpx

import "github.com/labstack/echo"

func RequestID(c echo.Context) string {
	if id := c.Response().Header().Get(echo.HeaderXRequestID); id != "" {
		return id
	}

	return c.Request().Header.Get(echo.HeaderXRequestID)
}
