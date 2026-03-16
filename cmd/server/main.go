package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/helioalb/finances/configs"
	"github.com/helioalb/finances/internal/user"
	"github.com/helioalb/finances/pkg/postgres"
	"github.com/labstack/echo"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	e := setupEcho()

	db := setupDatabase()
	defer db.Close()

	l := setupLogger()

	user.Init(e, db, l)

	l.Info("[server_starting][address=:8080]")
	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Error("[server_error][message=failed_to_start_server]", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	l.Info("[server_shutting_down]")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		l.Error("[server_error][message=failed_to_shutdown_server]", "error", err)
	}
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

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

func setupLogger() *slog.Logger {
	h := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(h)
}
