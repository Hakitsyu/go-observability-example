package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func ResponseTimeMetric(next http.Handler) http.Handler {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	histogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "api",
		Name:      "http_request_time_milliseconds",
		Buckets:   buckets,
	}, []string{"route", "method", "status_code"})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := NewStatusCodeRecorder(w)

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)

		route := getRoutePattern(r)
		method := r.Method
		statusCode := strconv.Itoa(rec.statusCode)

		histogram.WithLabelValues(route, method, statusCode).Observe(float64(elapsed.Milliseconds()))
	})
}

func getRoutePattern(r *http.Request) string {
	reqContext := chi.RouteContext(r.Context())
	if pattern := reqContext.RoutePattern(); pattern != "" {
		return pattern
	}

	return "null"
}
