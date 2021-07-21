package di

import (
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/controller"
)

func ProvideController(container *dig.Container) error {
	if err := container.Provide(controller.NewUserController); err != nil {
		return err
	}

	return nil
}
