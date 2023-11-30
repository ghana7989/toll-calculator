package main

import (
	"time"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.GPSData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"vehicle_id": data.UID,
			"latitude":   data.Lat,
			"longitude":  data.Lon,
			"took":       time.Since(start),
		}).Info("Received GPS data And Producing to Kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
