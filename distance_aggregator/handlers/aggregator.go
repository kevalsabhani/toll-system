package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Distance struct {
	Value     float64 `json:"value"`
	OBUId     int     `json:"obuId"`
	Timestamp int64   `json:"timestamp"`
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
		res := map[string]string{
			"error":   err.Error(),
			"message": "",
		}
		writeJsonResponse(w, http.StatusBadRequest, res)
		return
	}
	da.store.Insert(data)
	res := map[string]string{
		"error":   "",
		"message": "Distance added successfully",
	}
	writeJsonResponse(w, http.StatusOK, res)
}

func writeJsonResponse(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
