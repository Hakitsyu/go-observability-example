package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Hakitsyu/go-observability-example/api/handler"
	"github.com/Hakitsyu/go-observability-example/api/middleware"
	chi "github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	configureLogger()
	configureHttp()
}

func configureLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configureMiddlewareLogger()
}

func configureMiddlewareLogger() {
	chimiddleware.DefaultLogger = chimiddleware.RequestLogger(&chimiddleware.DefaultLogFormatter{Logger: log.Default(), NoColor: true})
}

func configureHttp() {
	router := chi.NewRouter()
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Logger)
	router.Use(middleware.ResponseTimeMetric)

	router.Get("/helloWorld", handler.HelloWorld)
	router.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, router)
}
