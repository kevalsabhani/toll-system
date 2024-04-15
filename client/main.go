/*
This client directory is an obu simulator. It generates
random coordinates and sends using websocket connection.
*/

package main

import (
	"context"
	"log"
	"time"

	"github.com/kevalsabhani/toll-calculator/client/obu"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var wsEndpoint = "ws://localhost:8080"

func main() {
	conn, _, err := websocket.Dial(context.Background(), wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.CloseNow()

	// Sending OBU data to server
	for {
		for i := 0; i < 20; i++ {
			obuData := obu.NewOBUData()
			log.Println("Sending OBU data: ", obuData)
			wsjson.Write(context.Background(), conn, obuData)
		}
		time.Sleep(time.Second)
	}
}
