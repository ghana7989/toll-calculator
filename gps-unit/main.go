package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/gorilla/websocket"
)

const wsEndPoint = "ws://localhost:8080/ws"

func sendGPSData(conn *websocket.Conn, data types.GPSData) error {
	return conn.WriteJSON(data)
}
func main() {
	uids := GenerateUID(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to server")
	}
	defer conn.Close()
	for {
		for _, uid := range uids {
			lat, long := GenerateLocation()
			data := types.GPSData{
				UID: uid,
				Lat: lat,
				Lon: long,
			}
			fmt.Printf("%+v\n", data)
			sendGPSData(conn, data)
		}
		time.Sleep(EmitterInterval)
	}
}
