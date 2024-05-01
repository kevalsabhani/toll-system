package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/client/services"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// OBUHandler contains connections and zap logger object
type OBUHandler struct {
	wsConn    *websocket.Conn
	kafkaConn *kafka.Conn
	logger    *zap.Logger
}

// NewOBUHandler return an object of OBUHandler
func NewOBUHandler(conn *kafka.Conn, logger *zap.Logger) *OBUHandler {
	return &OBUHandler{
		kafkaConn: conn,
		logger:    logger,
	}
}

// HandleWS handles websocket connection
func (h *OBUHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		h.logger.Fatal(err.Error())
	}

	h.wsConn = conn
	go func() {
		h.logger.Debug("New OBU client connected...")
		for {
			var data services.OBUData
			if err := wsjson.Read(context.Background(), h.wsConn, &data); err != nil {
				h.logger.Error(fmt.Sprintf("Read error: %s", err.Error()))
				continue
			}
			dataBytes, err := json.Marshal(data)
			if err != nil {
				h.logger.Error(fmt.Sprintf("Marshal error: %s", err.Error()))
				continue
			}

			_, err = h.kafkaConn.Write(dataBytes)
			if err != nil {
				h.logger.Error(err.Error())
			}
			h.logger.Info("OBU Data Received from client and published to kafka", zap.Int("obuId", data.OBUId))
		}
	}()
}
