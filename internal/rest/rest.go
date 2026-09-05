package rest

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/account"
	"github.com/helioalb/finances/internal/platform/httpx"
	"github.com/helioalb/finances/internal/transaction"
	"github.com/helioalb/finances/internal/user"
	"github.com/labstack/echo"
)

type handler struct {
	userSvc        user.Service
	accountSvc     account.Service
	transactionSvc transaction.Service
	log            echo.Logger
}

func NewHandler(userSvc user.Service, accountSvc account.Service, transactionSvc transaction.Service, log echo.Logger) *handler {
	return &handler{
		userSvc:        userSvc,
		accountSvc:     accountSvc,
		transactionSvc: transactionSvc,
		log:            log,
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.POST("/users", func(c echo.Context) error {
		var input user.CreateInput
		requestID := httpx.RequestID(c)

		if err := c.Bind(&input); err != nil {
			return h.badRequestResponse(c, err)
		}

		if err := input.Validate(); err != nil {
			return h.unprocessableEntityResponse(c, err)
		}

		ctx := c.Request().Context()

		u, err := h.userSvc.Create(ctx, input)
		if err != nil {
			if errors.Is(err, user.ErrEmailInUse) {
				return h.emailAlreadyInUseResponse(c)
			}

			return h.internalServerErrorResponse(c, err)
		}

		c.Logger().Info(
			"[user][create]",
			"[http_status=", http.StatusCreated, "]",
			"[user_uuid=", u.UUID.String(), "]",
			"[request_id=", requestID, "]",
		)

		return c.JSON(http.StatusCreated, map[string]string{
			"uuid": u.UUID.String(),
		})
	})

	e.POST("/accounts", func(c echo.Context) error {
		var input account.CreateInput

		if err := c.Bind(&input); err != nil {
			return h.badRequestResponse(c, err)
		}

		if err := input.Validate(); err != nil {
			return h.unprocessableEntityResponse(c, err)
		}

		ctx := c.Request().Context()

		a, err := h.accountSvc.Create(ctx, input)
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				return h.userNotFoundResponse(c)
			}

			if errors.Is(err, account.ErrAccountAlreadyExists) {
				return h.accountAlreadyExistsResponse(c)
			}

			return h.internalServerErrorResponse(c, err)
		}

		h.log.Info(
			"[account][create]",
			"[http_status=", http.StatusCreated, "]",
			"[account_uuid=", a.UUID.String(), "]",
		)

		return c.JSON(http.StatusCreated, map[string]string{
			"uuid": a.UUID.String(),
		})
	})

	e.POST("/transactions/expense", func(c echo.Context) error {
		var input transaction.CreateInput

		if err := c.Bind(&input); err != nil {
			return h.badRequestResponse(c, err)
		}

		if err := input.Validate(); err != nil {
			return h.unprocessableEntityResponse(c, err)
		}

		ctx := c.Request().Context()

		accountUUID, err := uuid.Parse(input.AccountUUID)
		if err != nil {
			return h.badRequestResponse(c, err)
		}

		err = h.transactionSvc.Expense(ctx, accountUUID, input.Amount, input.Description)
		if err != nil {
			if errors.Is(err, transaction.ErrAccountNotFound) {
				return h.accountNotFoundResponse(c)
			}

			return h.internalServerErrorResponse(c, err)
		}

		return c.NoContent(http.StatusCreated)
	})

	e.POST("/transactions/income", func(c echo.Context) error {
		var input transaction.CreateInput

		if err := c.Bind(&input); err != nil {
			return h.badRequestResponse(c, err)
		}

		if err := input.Validate(); err != nil {
			return h.unprocessableEntityResponse(c, err)
		}

		ctx := c.Request().Context()

		accountUUID, err := uuid.Parse(input.AccountUUID)
		if err != nil {
			return h.badRequestResponse(c, err)
		}

		err = h.transactionSvc.Income(ctx, accountUUID, input.Amount, input.Description)
		if err != nil {
			if errors.Is(err, transaction.ErrAccountNotFound) {
				return h.accountNotFoundResponse(c)
			}

			return h.internalServerErrorResponse(c, err)
		}

		return c.NoContent(http.StatusCreated)
	})

	e.POST("/transactions/transfer", func(c echo.Context) error {
		return nil
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

func (h *handler) accountAlreadyExistsResponse(c echo.Context) error {
	requestID := httpx.RequestID(c)

	h.log.Warn(
		"[account][create]",
		"[http_status=", http.StatusConflict, "]",
		"[error=account already exists]",
		"[request_id=", requestID, "]",
	)

	return c.JSON(http.StatusConflict, map[string]string{
		"error": "account already exists",
	})
}

func (h *handler) accountNotFoundResponse(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "account not found",
	})
}

func (h *handler) userNotFoundResponse(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "user not found",
	})
}
