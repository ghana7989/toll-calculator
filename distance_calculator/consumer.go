package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ghana7989/toll-calculator/aggregator/client"
	"github.com/ghana7989/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
	aggClient   *client.Client
}

func NewKafkaConsumer(topic string, svc CalculatorServicer, client *client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "distance-calculator",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient:   client,
	}, nil
}
func (c *KafkaConsumer) Start() {
	logrus.Info("Starting consumer")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Consumer error: %v (%v)\n", err, msg)
			continue
		} else if err == nil {
			data := &types.GPSData{}
			err := json.Unmarshal(msg.Value, data)
			if err != nil {
				logrus.Errorf("Error unmarshalling data: %v", err)
				continue
			}
			distance, err := c.calcService.CalculateDistance(*data)
			if err != nil {
				logrus.Errorf("Calculation Error %s", err)
				continue
			}
			distanceData := types.Distance{
				UID:   data.UID,
				Value: distance,
				Unix:  time.Now().UnixNano(),
			}
			err = c.aggClient.AggregateInvoice(distanceData)
			if err != nil {
				logrus.Errorf("Error aggregating invoice: %v", err)
				continue
			}
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
