package transaction

import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo, transaction *handler) {
	e.POST("/transactions/expense", transaction.Expense)
	e.POST("/transactions/income", transaction.Income)
	e.POST("/transactions/transfer", transaction.Transfer)
}
