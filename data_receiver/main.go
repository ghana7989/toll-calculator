package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch    chan types.GPSData
	conn     *websocket.Conn
	producer DataProducer
}

func main() {
	var (
		p   DataProducer
		err error
	)
	p, err = NewKafkaProducer("gps-data")
	if err != nil {
		log.Fatal("Failed to set up Kafka producer:", err)
	}
	p = NewLogMiddleware(p)

	msgch := make(chan types.GPSData, 128)
	dr := &DataReceiver{
		msgch:    msgch,
		producer: p,
	}

	http.HandleFunc("/ws", dr.wsHandler)
	go handleIncomingMessages(msgch)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIncomingMessages(msg chan types.GPSData) {
	for data := range msg {
		// Do something with the data
		fmt.Printf("%+v\n", data)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (dr *DataReceiver) produceData(data *types.GPSData) error {
	return dr.producer.ProduceData(*data)
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
		err = dr.produceData(data)
		if err != nil {
			log.Println("Failed to produce data", err)
			return
		}
		// dr.msgch <- *data
	}
}
