package di

import (
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/common"
)

func ProvideConfig(container *dig.Container) error {
	if err := container.Provide(common.ConfigFromEnv); err != nil {
		return err
	}

	return nil
}
