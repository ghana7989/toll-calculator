package main

import (
	"time"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}
func (l *LogMiddleware) AggregateDistance(distance types.Distance) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"distance": distance,
			"took":     time.Since(start),
		}).Info("Aggregating distance")
	}(time.Now())
	return l.next.AggregateDistance(distance)
}
