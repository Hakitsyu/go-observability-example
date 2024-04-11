package prometheus

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ConfigureMetricsRoute(mux *chi.Mux, pattern string) {
	mux.Handle("/metrics", promhttp.Handler())
}
