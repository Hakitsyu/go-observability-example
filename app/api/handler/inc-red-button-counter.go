package handler

import (
	"log"
	"net/http"

	"github.com/Hakitsyu/go-observability-example/internal/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter   = otel.Meter("user")
	counter metric.Int64Counter
)

func init() {
	var err error
	counter, err = meter.Int64Counter("index_red_button_click")

	if err != nil {
		log.Fatal(err)
	}
}

func IncRedButtonCounter(writer *api.ResponseWriter, request *http.Request) {
	counter.Add(request.Context(), 1)
}
