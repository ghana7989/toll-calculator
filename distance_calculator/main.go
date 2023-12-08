package main

import (
	"github.com/ghana7989/toll-calculator/aggregator/client"
	"github.com/sirupsen/logrus"
)

const (
	aggregateEndpoint = "http://localhost:3000/aggregate"
)

func main() {
	calcService := NewCalculatorService()
	calcService = NewLogMiddleware(calcService)
	client := client.NewClient(aggregateEndpoint)
	kafkaConsumer, err := NewKafkaConsumer("gps-data", calcService, client)
	if err != nil {
		logrus.Fatalf("Error creating consumer: %v", err)
	}
	kafkaConsumer.Start()
}
