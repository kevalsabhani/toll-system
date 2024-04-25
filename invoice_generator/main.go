package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/handlers"
	"go.uber.org/zap"
)

const (
	port = ":3000"
	env  = "DEVELOPMENT"
)

var logger *zap.Logger

func main() {
	//zap logger
	logger = zap.Must(zap.NewProduction())
	if env == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	store := handlers.NewMemoryStore(logger)
	aggregator := handlers.NewDistanceAggregator(store, logger)
	generator := handlers.NewInvoiceGenerator(store, logger)

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go makeHttpTransport(aggregator, generator)

	<-sigchan
	logger.Info("Closing the server...")
}

func makeHttpTransport(aggregator *handlers.DistanceAggregator, generator *handlers.InvoiceGenerator) {
	logger.Info("Http transport running on", zap.String("port", port))
	http.HandleFunc("/aggregate_distance", aggregator.AggregateDistance)
	http.HandleFunc("/invoice", generator.GenerateInvoice)
	http.ListenAndServe(port, nil)
}
