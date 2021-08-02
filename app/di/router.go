package di

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/controller"
	"mayerkv/otus-ma-lesson-6/app/metrics"
	"net/http"
)

func createRouter(userController *controller.UserController, r *prometheus.Registry) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(r, promhttp.HandlerOpts{})))

	users := router.Group("/user")
	users.Use(metrics.MetricsMiddleWare)

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
