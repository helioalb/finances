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
	e := echo.New()
	db := setupDatabase()
	logger := setupLogger()

	_ = user.Init(e, db, logger)

	logger.Info("[server_starting][address=:8080]")
	e.Start(":8080")
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func setupDatabase() *sql.DB {
	cfg := configs.PostgresConfig()

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
