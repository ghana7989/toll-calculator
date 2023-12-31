package main

import "github.com/ghana7989/toll-calculator/types"

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m MemoryStore) Insert(d types.Distance) error {
	m.data[d.UID] += d.Value
	return nil
}
