package di

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
	"mayerkv/otus-ma-lesson-6/app/metrics"
)

func ProvideMetrics(container *dig.Container) error {
	if err := container.Provide(createRegistrar); err != nil {
		return err
	}

	return nil
}

func createRegistrar() (*prometheus.Registry, error) {
	r := prometheus.NewRegistry()

	if err := r.Register(metrics.RequestCount); err != nil {
		return nil, err
	}
	if err := r.Register(metrics.RequestLatency); err != nil {
		return nil, err
	}

	return r, nil
}
