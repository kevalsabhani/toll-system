package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kevalsabhani/toll-calculator/client/obu"
	"github.com/segmentio/kafka-go"
)

const (
	env       = "DEVELOPMENT"
	kafkaHost = "localhost:9092"
	groupId   = "distance-calculator"
	topic     = "obu-data"
)

var coords = [2][2]float64{}

func main() {
	//zap logger
	// logger := zap.Must(zap.NewProduction())
	// if env == "DEVELOPMENT" {
	// 	logger = zap.Must(zap.NewDevelopment())
	// }

	// Set up a Kafka reader
	config := kafka.ReaderConfig{
		Brokers:  []string{kafkaHost},
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}
	reader := kafka.NewReader(config)
	defer reader.Close()

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Loop to read messages
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ReadMessage blocks until it receives a message or ctx times out
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		// Process the message
		// fmt.Printf("Received message: %s\n", msg.Value)
		distance := distanceCal(msg.Value)
		fmt.Printf("Distance travelled: %f\n", distance)
		select {
		case <-sigchan:
			fmt.Println("Received termination signal. Closing consumer.")
			return
		default:
		}
	}
}

func distanceCal(obuData []byte) float64 {
	var data obu.OBUData
	if err := json.Unmarshal(obuData, &data); err != nil {
		log.Println("Unmarshal error: ", err)
	}
	if coords[1][0] == 0 && coords[1][1] == 0 {
		coords[1][0] = data.Lat
		coords[1][1] = data.Long
		return 0
	}
	coords[0][0] = coords[1][0]
	coords[0][1] = coords[1][1]
	coords[1][0] = data.Lat
	coords[1][1] = data.Long
	return math.Sqrt(math.Pow((coords[1][0]-coords[0][0]), 2) + math.Pow((coords[1][1]-coords[0][1]), 2))
}
