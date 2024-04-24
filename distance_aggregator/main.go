package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevalsabhani/toll-calculator/distance_aggregator/handlers"
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

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go makeHttpTransport(aggregator)

	<-sigchan
	logger.Info("Closing the server...")
}

func makeHttpTransport(aggregator *handlers.DistanceAggregator) {
	logger.Info("Http transport running on", zap.String("port", port))
	http.HandleFunc("/aggregate", aggregator.AggregateDistance)
	http.ListenAndServe(port, nil)
}
