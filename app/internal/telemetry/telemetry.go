package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

type ConfigureTelemetryHandleError func(error)

func ConfigureTelemetry(handleError ConfigureTelemetryHandleError) {
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleError(err)
		return
	}

	otel.SetMeterProvider(meterProvider)
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
