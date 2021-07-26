package common

import (
	"os"
	"strconv"
)

type Config struct {
	DbDsn         string
	GracefulDelay int
}

func NewConfig(dbDsn string, d int) *Config {
	return &Config{
		DbDsn:         dbDsn,
		GracefulDelay: d,
	}
}

func ConfigFromEnv() *Config {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgresql://postgres:@localhost:5432/users"
	}
	delay, err := strconv.Atoi(os.Getenv("GRACEFUL_DELAY"))
	if err != nil {
		delay = 5
	}

	return NewConfig(dsn, delay)
}
