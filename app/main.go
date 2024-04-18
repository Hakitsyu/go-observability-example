package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/Hakitsyu/go-observability-example/api/handler"
	"github.com/Hakitsyu/go-observability-example/api/middleware"
	"github.com/Hakitsyu/go-observability-example/internal/api"
	"github.com/Hakitsyu/go-observability-example/internal/telemetry"
	"github.com/Hakitsyu/go-observability-example/internal/telemetry/prometheus"
	chi "github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()

	configureLogger()

	if err := prometheus.UseMetricProvider(); err != nil {
		log.Fatal(err)
		return
	}

	if err := telemetry.ConfigureOpenTelemetry(ctx); err != nil {
		log.Fatal(err)
		return
	}

	configureHttp(ctx)
}

func configureLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configureMiddlewareLogger()
}

func configureMiddlewareLogger() {
	chimiddleware.DefaultLogger = chimiddleware.RequestLogger(&chimiddleware.DefaultLogFormatter{Logger: log.Default(), NoColor: true})
}

func configureHttp(ctx context.Context) {
	router := chi.NewRouter()
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Logger)
	router.Use(middleware.ResponseTimeMetric)

	router.Get("/helloWorld", newDefaultMuxHandler(handler.HelloWorld))
	router.Get("/", newDefaultMuxHandler(handler.Index))
	router.Post("/inc-red-button-counter", newDefaultMuxHandler(handler.IncRedButtonCounter))

	prometheus.ConfigureMetricsRoute(router, "/metrics")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := &http.Server{
		Addr:        ":" + port,
		BaseContext: func(l net.Listener) context.Context { return ctx },
		Handler:     router,
	}

	app.ListenAndServe()
}

func newDefaultMuxHandler(handler func(writer *api.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		handler(api.NewResponseWriter(w), request)
	}
}
