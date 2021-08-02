package metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "app_request_count",
		Help: "Application Request Count",
	}, []string{"method", "endpoint", "http_status"})

	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "app_request_latency_seconds",
		Help:    "Application Request Latency",
		Buckets: nil,
	}, []string{"method", "endpoint"})
)

func MetricsMiddleWare(c *gin.Context) {
	method := c.Request.Method
	endpoint := c.FullPath()
	timer := prometheus.NewTimer(RequestLatency.WithLabelValues(method, endpoint))

	c.Next()
	status := fmt.Sprint(c.Writer.Status())

	RequestCount.WithLabelValues(method, endpoint, status).Inc()
	timer.ObserveDuration()
}
