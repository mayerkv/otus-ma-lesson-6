package common

import (
	"fmt"
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
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if dsn == "" {
		dsn = "postgresql://postgres:@localhost:5432/users"
	}
	delay, err := strconv.Atoi(os.Getenv("GRACEFUL_DELAY"))
	if err != nil {
		delay = 5
	}

	return NewConfig(dsn, delay)
}
