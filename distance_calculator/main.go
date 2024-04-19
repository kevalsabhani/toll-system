package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevalsabhani/toll-calculator/distance_calculator/services"
	"go.uber.org/zap"
)

const (
	env = "DEVELOPMENT"
)

func main() {
	//zap logger
	logger := zap.Must(zap.NewProduction())
	if env == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer := services.NewKafkaConsumer(logger)

	// Loop to read messages
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		data := consumer.Read(ctx)

		// Calculate the distance
		distance := services.DistanceCal(*data)
		logger.Info("Distance travelled:", zap.Float64("distance", distance))

		select {
		case <-sigchan:
			logger.Info("Closing consumer...")
			return
		default:
		}
	}
}
