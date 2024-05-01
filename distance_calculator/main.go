package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevalsabhani/toll-calculator/distance_calculator/services"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/client"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/types"
	"github.com/kevalsabhani/toll-calculator/pb"
	"go.uber.org/zap"
)

const (
	env                      = "DEVELOPMENT"
	invoiceGeneratorHttpHost = "http://localhost:3000"
	invoiceGeneratorGrpcHost = "localhost:50051"
)

func main() {
	transportPtr := flag.String("t", "http", "select the transport[http | grcp]")
	flag.Parse()

	//zap logger
	logger := zap.Must(zap.NewProduction())
	if env == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer := services.NewKafkaConsumer(logger)

	var (
		invoiceGeneratorHttpClient *client.InvoiceGeneratorClient
		invoiceGeneratorGrpcClient pb.InvoiceServiceClient
		err                        error
	)
	if *transportPtr == "http" {
		invoiceGeneratorHttpClient = client.NewInvoiceGeneratorClient(
			fmt.Sprintf("%s/aggregate_distance", invoiceGeneratorHttpHost),
		)
	} else {
		invoiceGeneratorGrpcClient, err = client.NewInvoiceGeneratorGrpcClient(invoiceGeneratorGrpcHost)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	// Loop to read messages
	for {
		data := consumer.Read(context.Background())

		// Calculate the distance
		distance := services.DistanceCal(*data)

		// Post Distance data to distance aggregator service
		distanceData := types.Distance{
			Value:     distance,
			OBUId:     data.OBUId,
			Timestamp: time.Now().UnixNano(),
		}

		if *transportPtr == "http" {
			if err := invoiceGeneratorHttpClient.PostDistanceData(&distanceData); err != nil {
				logger.Error(
					fmt.Sprintf("failed to post distance for OBUId: %d", data.OBUId),
				)
			}
		} else {
			if _, err := invoiceGeneratorGrpcClient.AggregateDistance(context.Background(), &pb.AggregateDistanceRequest{
				Value:     distanceData.Value,
				OBUId:     int64(distanceData.OBUId),
				Timestamp: distanceData.Timestamp,
			}); err != nil {
				logger.Error(err.Error())
				continue
			}
		}

		logger.Info(
			"distance data read from kafka and posted to invoice service successfully...",
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
