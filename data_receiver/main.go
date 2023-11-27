package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/gorilla/websocket"
)

func main() {
	msg := make(chan types.GPSData)
	dr := &DataReceiver{
		msg: msg,
	}
	http.HandleFunc("/ws", dr.wsHandler)
	go func() {
		for range msg {
			// Do something with the data
			fmt.Printf("%+v\n", <-msg)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type DataReceiver struct {
	msg  chan types.GPSData
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (dr *DataReceiver) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Failed to upgrade to websocket")
		return
	}
	dr.conn = conn
	defer conn.Close()
	for {
		data := &types.GPSData{}
		err := conn.ReadJSON(data)
		if err != nil {
			log.Println("Failed to read json", err)
			return
		}
		dr.msg <- *data
	}
}
