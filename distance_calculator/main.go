package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevalsabhani/toll-calculator/distance_calculator/services"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/client"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/types"
	"go.uber.org/zap"
)

const (
	env                  = "DEVELOPMENT"
	invoiceGeneratorHost = "http://localhost:3000"
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

	aggregatorClient := client.NewInvoiceGeneratorClient(
		fmt.Sprintf("%s/aggregate_distance", invoiceGeneratorHost),
	)

	// Loop to read messages
	for {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		data := consumer.Read(context.Background())

		// Calculate the distance
		distance := services.DistanceCal(*data)

		// Post Distance data to distance aggregator service
		distanceData := types.Distance{
			Value:     distance,
			OBUId:     data.OBUId,
			Timestamp: time.Now().UnixNano(),
		}
		if err := aggregatorClient.PostDistanceData(&distanceData); err != nil {
			logger.Error(
				fmt.Sprintf("failed to post distance for OBUId: %d", data.OBUId),
			)
		}
		logger.Info(
			"distance data posted successfully...",
			zap.Float64("distance", distanceData.Value),
			zap.Int("obuId", distanceData.OBUId),
		)

		select {
		case <-sigchan:
			logger.Info("Closing consumer...")
			return
		default:
		}
	}
}
