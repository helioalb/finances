package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	"github.com/helioalb/finances/configs"
	"github.com/helioalb/finances/internal/user"
	"github.com/helioalb/finances/pkg/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := setupEcho()
	db := setupDatabase()
	l := setupLogger()

	user.Init(e, db, l)

	l.Info("[server_starting][address=:8080]")
	e.Start(":8080")
}

func setupEcho() *echo.Echo {
	e := echo.New()

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
