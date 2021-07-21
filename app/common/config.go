package common

import "os"

type Config struct {
	DbDsn string
	Addr  string
}

func NewConfig(dbDsn, addr string) *Config {
	return &Config{
		DbDsn: dbDsn,
		Addr:  addr,
	}
}

func ConfigFromEnv() *Config {
	return NewConfig(os.Getenv("DB_DSN"), os.Getenv("ADDR"))
}
