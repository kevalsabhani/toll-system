package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Distance struct {
	Value float64 `json:"value"`
	OBUId int     `json:"obuId"`
}

type DistanceAggregator struct {
	store  Store
	logger *zap.Logger
}

func NewDistanceAggregator(store Store, logger *zap.Logger) *DistanceAggregator {
	return &DistanceAggregator{
		store:  store,
		logger: logger,
	}
}

func (da *DistanceAggregator) AggregateDistance(w http.ResponseWriter, r *http.Request) {
	data := new(Distance)
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		da.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	da.store.Insert(data)
}
