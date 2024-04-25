package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/utils"
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
		utils.WriteJsonResponse(w, http.StatusBadRequest, res)
		return
	}
	da.store.Insert(data)
	res := map[string]string{
		"error":   "",
		"message": "Distance added successfully",
	}
	utils.WriteJsonResponse(w, http.StatusOK, res)
}
