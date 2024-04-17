package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	ENV       = "DEVELOPMENT"
	kafkaHost = "localhost:9092"
)

func main() {
	//zap logger
	// logger := zap.Must(zap.NewProduction())
	// if ENV == "DEVELOPMENT" {
	// 	logger = zap.Must(zap.NewDevelopment())
	// }

	topic := "obu-data"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
