package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func Connect(cfg Config) (*pgxpool.Pool, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid database configuration: %w", err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	if cfg.MaxOpenConns > 0 {
		poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns > 0 {
		poolConfig.MinConns = int32(cfg.MaxIdleConns)
		if poolConfig.MaxConns > 0 && poolConfig.MinConns > poolConfig.MaxConns {
			poolConfig.MinConns = poolConfig.MaxConns
		}
	}

	if cfg.ConnMaxLifetime > 0 {
		poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
