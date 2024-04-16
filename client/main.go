/*
This client directory is an obu simulator. It generates
random coordinates and sends using websocket connection.
*/

package main

import (
	"context"
	"time"

	"github.com/kevalsabhani/toll-calculator/client/obu"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const (
	ENV        = "DEVELOPMENT"
	wsEndpoint = "ws://localhost:8080"
)

func main() {

	//zap logger
	logger := zap.Must(zap.NewProduction())
	if ENV == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	// websocket connection
	conn, _, err := websocket.Dial(context.Background(), wsEndpoint, nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.CloseNow()

	// Sending OBU data to server
	for {
		for i := 0; i < 20; i++ {
			obuData := obu.NewOBUData()
			logger.Info("Sending OBU data", zap.Any("obu", obuData))
			wsjson.Write(context.Background(), conn, obuData)
		}
		time.Sleep(5 * time.Second)
	}
}
