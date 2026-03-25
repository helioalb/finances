package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/helioalb/finances/configs"
	"github.com/helioalb/finances/internal/account"
	"github.com/helioalb/finances/internal/transaction"
	"github.com/helioalb/finances/internal/user"
	"github.com/helioalb/finances/pkg/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	e := setupEcho()
	log := e.Logger

	db := setupDatabase()
	defer db.Close()

	userSvc := user.Init(e, db, log)
	accountSvc := account.Init(e, db, userSvc, log)
	transaction.Init(e, db, accountSvc)

	log.Info("[server_starting][address=:8080]")
	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("[server_error][message=failed_to_start_server]", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Info("[server_shutting_down]")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Error("[server_error][message=failed_to_shutdown_server]", "error", err)
	}
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetPrefix("finances")
	e.Logger.SetHeader(`{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}","message":"${message}"}`)

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	return e
}

func setupDatabase() *sql.DB {
	cfg := configs.PostgresConfig()

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf(
			"Failed to connect to database: %v",
			err,
		)
	}

	return db
}
