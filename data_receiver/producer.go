package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ghana7989/toll-calculator/types"
)

type DataProducer interface {
	ProduceData(types.GPSData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    *string
}

func NewKafkaProducer(topic string) (DataProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: p, topic: &topic}, nil
}

func (p *KafkaProducer) ProduceData(data types.GPSData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)
}
