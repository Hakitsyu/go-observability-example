package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	TraceProvider *trace.TracerProvider
	MeterProvider *metric.MeterProvider
)

func ConfigureOpenTelemetry(ctx context.Context) error {
	if err := setupMeterProvider(ctx); err != nil {
		return err
	}

	otel.SetMeterProvider(MeterProvider)
	return nil
}

func setupMeterProvider(ctx context.Context) error {
	var err error
	if MeterProvider != nil {
		return nil
	}

	MeterProvider, err = newMeterProvider(ctx)
	return err
}

func newMeterProvider(ctx context.Context) (*metric.MeterProvider, error) {
	exporter, err := newMeterExporter(ctx)
	if err != nil {
		return nil, err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
	)

	return provider, nil
}

func newMeterExporter(ctx context.Context) (metric.Exporter, error) {
	exporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	return exporter, nil
}
