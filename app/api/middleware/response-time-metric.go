package middleware

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter     = otel.Meter("http")
	histogram metric.Float64Histogram
)

func init() {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	var err error
	histogram, err = meter.Float64Histogram("http_request_time_milliseconds",
		metric.WithExplicitBucketBoundaries(buckets...))

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseTimeMetric(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := NewStatusCodeRecorder(w)

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)

		route := getRoutePattern(r)
		method := r.Method
		statusCode := strconv.Itoa(rec.statusCode)

		attributes := []attribute.KeyValue{
			attribute.String("route", route),
			attribute.String("method", method),
			attribute.String("statusCode", statusCode),
		}

		histogram.Record(r.Context(), float64(elapsed.Milliseconds()), metric.WithAttributes(attributes...))
	})
}

func getRoutePattern(r *http.Request) string {
	reqContext := chi.RouteContext(r.Context())
	if pattern := reqContext.RoutePattern(); pattern != "" {
		return pattern
	}

	return "null"
}
