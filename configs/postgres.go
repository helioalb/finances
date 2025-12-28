package configs

import "github.com/helioalb/finances/pkg/postgres"

func PostgresConfig() postgres.Config {
	return postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "finances",
		Password: "finances",
		DBName:   "finances",
		SSLMode:  "disable",
	}
}
