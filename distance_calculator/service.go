package main

import (
	"math"

	"github.com/ghana7989/toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.GPSData) (float64, error)
}

type CalculatorService struct {
	prevPoint    types.GPSData
	currentPoint types.GPSData
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		prevPoint:    types.GPSData{},
		currentPoint: types.GPSData{},
	}
}

func (s *CalculatorService) CalculateDistance(data types.GPSData) (float64, error) {
	if s.prevPoint.Lat == 0.0 && s.prevPoint.Lon == 0.0 {
		s.prevPoint = data
		return 0.0, nil
	}
	s.currentPoint = data
	distance := calculateDistance(s.prevPoint.Lat, s.prevPoint.Lon, s.currentPoint.Lat, s.currentPoint.Lon)
	s.prevPoint = s.currentPoint
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
