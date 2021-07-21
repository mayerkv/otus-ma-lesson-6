package di

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/controller"
)

func createRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	users := router.Group("/user")
	for _, h := range userController.Handlers() {
		users.Handle(h.Method, h.Path, h.HandlerFunc)
	}

	return router
}

func ProvideRouter(container *dig.Container) error {
	if err := container.Provide(createRouter); err != nil {
		return err
	}

	return nil
}
