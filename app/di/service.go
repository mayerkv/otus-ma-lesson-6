package di

import (
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/domain"
)

func ProvideService(container *dig.Container) error {
	if err := container.Provide(domain.NewUserService); err != nil {
		return err
	}

	return nil
}
