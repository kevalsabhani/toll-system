package handlers

import (
	"net/http"
	"strconv"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/utils"
	"go.uber.org/zap"
)

type Invoice struct {
	OBUId         int     `json:"obuId"`
	TotalDistance float64 `json:"totalDistance"`
	Amount        float64 `json:"amount"`
}

type InvoiceGenerator struct {
	store  Store
	logger *zap.Logger
}

func NewInvoiceGenerator(store Store, logger *zap.Logger) *InvoiceGenerator {
	return &InvoiceGenerator{
		store:  store,
		logger: logger,
	}
}

func (ig *InvoiceGenerator) GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	obuId, err := strconv.Atoi(r.URL.Query().Get("obuid"))
	if err != nil {
		utils.WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{
				"message": "",
				"error":   err.Error(),
			},
		)
		return
	}
	totalDistance, err := ig.store.Get(obuId)
	if err != nil {
		utils.WriteJsonResponse(
			w,
			http.StatusBadRequest,
			map[string]string{
				"message": "",
				"error":   err.Error(),
			},
		)
		return
	}
	invoice := Invoice{
		OBUId:         obuId,
		TotalDistance: totalDistance,
		Amount:        totalDistance * 10.0,
	}
	utils.WriteJsonResponse(w, http.StatusOK, invoice)
}
