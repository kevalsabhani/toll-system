package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kevalsabhani/toll-calculator/server/handlers"
	"github.com/segmentio/kafka-go"
)

func main() {

	// configure kafka produucer
	topic := "obu-data"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }

	obuHandler := handlers.NewOBUHandler(conn)
	http.HandleFunc("/", obuHandler.HandleWS)
	http.ListenAndServe(":8080", nil)
}
