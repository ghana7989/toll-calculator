package main

import (
	"math"
	"math/rand"
	"time"
)

const EmitterInterval = 1 * time.Second

func GenerateLocation() (float64, float64) {
	return GenerateCoordinate(), GenerateCoordinate()
}
func GenerateCoordinate() float64 {
	return rand.Float64()*180 - 90
}
func GenerateUID(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt64)
	}
	return ids
}
