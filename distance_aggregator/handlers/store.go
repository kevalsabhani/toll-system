package handlers

import "go.uber.org/zap"

type Store interface {
	Insert(*Distance)
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
	ms.logger.Info("Distance added to memory", zap.Float64("distance", d.Value))
}
