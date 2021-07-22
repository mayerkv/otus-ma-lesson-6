package di

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/common"
	"time"
)

func ProvideDatabase(container *dig.Container) error {
	if err := container.Provide(createDb); err != nil {
		return err
	}

	return nil
}

func createDb(config *common.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", config.DbDsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetConnMaxIdleTime(60 * time.Second)

	return db, nil
}
