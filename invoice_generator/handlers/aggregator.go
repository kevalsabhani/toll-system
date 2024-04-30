package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/storage"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/types"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/utils"
	"go.uber.org/zap"
)

type DistanceAggregator struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDistanceAggregator(store storage.Store, logger *zap.Logger) *DistanceAggregator {
	return &DistanceAggregator{
		store:  store,
		logger: logger,
	}
}

func (da *DistanceAggregator) AggregateDistance(w http.ResponseWriter, r *http.Request) {
	data := new(types.Distance)
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
