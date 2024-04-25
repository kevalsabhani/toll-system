package handlers

import (
	"fmt"

	"go.uber.org/zap"
)

type Store interface {
	Insert(*Distance)
	Get(int) (float64, error)
}

type MemoryStore struct {
	DistanceData map[int]float64
	logger       *zap.Logger
}

func NewMemoryStore(logger *zap.Logger) *MemoryStore {
	return &MemoryStore{
		DistanceData: make(map[int]float64),
		logger:       logger,
	}
}

func (ms *MemoryStore) Insert(d *Distance) {
	ms.DistanceData[d.OBUId] += d.Value
	ms.logger.Info(
		"distance data stored successfully...",
		zap.Int("obuId", d.OBUId),
		zap.Float64("distance", d.Value),
	)
}

func (ms *MemoryStore) Get(obuId int) (float64, error) {
	totalDistance, ok := ms.DistanceData[obuId]
	if !ok {
		return 0.0, fmt.Errorf("OBUId [%d] not found", obuId)
	}
	return totalDistance, nil
}
