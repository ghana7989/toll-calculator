package main

import "github.com/sirupsen/logrus"

func main() {
	calcService := NewCalculatorService()
	calcService = NewLogMiddleware(calcService)
	kafkaConsumer, err := NewKafkaConsumer("gps-data", calcService)
	if err != nil {
		logrus.Fatalf("Error creating consumer: %v", err)
	}
	kafkaConsumer.Start()
}
