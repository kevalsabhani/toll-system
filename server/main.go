package main

import (
	"context"
	"log"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/server/handlers"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const ENV = "DEVELOPMENT"

func main() {
	//zap logger
	logger := zap.Must(zap.NewProduction())
	if ENV == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	// kafka producer
	kafkaConn, err := configureKafkaProducer()
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	logger.Debug("Kafka connected...")

	obuHandler := handlers.NewOBUHandler(kafkaConn, logger)

	//routes
	http.HandleFunc("/", obuHandler.HandleWS)
	logger.Info("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// configureKafkaProducer returns kafka connection object
func configureKafkaProducer() (*kafka.Conn, error) {
	topic := "obu-data"
	partition := 0
	return kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
}
