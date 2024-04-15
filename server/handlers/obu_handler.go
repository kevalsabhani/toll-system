package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/client/obu"
	"github.com/segmentio/kafka-go"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type OBUHandler struct {
	wsConn    *websocket.Conn
	kafkaConn *kafka.Conn
}

func NewOBUHandler(conn *kafka.Conn) *OBUHandler {
	return &OBUHandler{
		kafkaConn: conn,
	}
}

func (h *OBUHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	h.wsConn = conn
	go func() {
		log.Println("New OBU client connected...")
		for {
			var data obu.OBUData
			if err := wsjson.Read(context.Background(), h.wsConn, &data); err != nil {
				log.Println("Read error: ", err)
				continue
			}
			log.Println("Received OBU data: ", data)
			dataBytes, err := json.Marshal(data)
			if err != nil {
				log.Println("Failed to marshal obu data: ", err)
				continue
			}

			_, err = h.kafkaConn.Write(dataBytes)
			if err != nil {
				log.Println("Failed to write message:", err)
			}
		}
	}()
}
