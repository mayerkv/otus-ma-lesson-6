package di

import (
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/infrastructure/persistence/pg"
)

func ProvideRepository(container *dig.Container) error {
	if err := container.Provide(pg.NewUserRepository); err != nil {
		return err
	}

	return nil
}
