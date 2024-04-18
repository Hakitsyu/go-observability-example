package prometheus

import (
	"github.com/Hakitsyu/go-observability-example/internal/telemetry"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func UseMetricProvider() error {
	provider, err := newMeterProvider()
	if err != nil {
		return err
	}

	telemetry.MeterProvider = provider
	return nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	exporter, err := newMeterExporter()
	if err != nil {
		return nil, err
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))

	return provider, nil
}

func newMeterExporter() (metric.Reader, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	return exporter, nil
}

func ConfigureMetricsRoute(mux *chi.Mux, pattern string) {
	mux.Handle("/metrics", promhttp.Handler())
}
