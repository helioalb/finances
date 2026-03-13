package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func (c Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}

	if c.Port == "" {
		return fmt.Errorf("port is required")
	}

	if c.User == "" {
		return fmt.Errorf("user is required")
	}

	if c.Password == "" {
		return fmt.Errorf("password is required")
	}

	if c.DBName == "" {
		return fmt.Errorf("dbname is required")
	}

	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}

	if c.MaxOpenConns < 0 {
		return fmt.Errorf("max open connections cannot be negative")
	}

	if c.MaxIdleConns < 0 {
		return fmt.Errorf("max idle connections cannot be negative")
	}

	if c.ConnMaxLifetime < 0 {
		return fmt.Errorf("connection max lifetime cannot be negative")
	}

	return nil
}

func Connect(cfg Config) (*sql.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid database configuration: %w", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		if err := db.Close(); err != nil {
			return nil, fmt.Errorf("failed to close database: %w", err)
		}
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
