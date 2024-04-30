package handlers

import (
	"net/http"
	"strconv"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/storage"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/types"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/utils"
	"go.uber.org/zap"
)

type InvoiceGenerator struct {
	store  storage.Store
	logger *zap.Logger
}

func NewInvoiceGenerator(store storage.Store, logger *zap.Logger) *InvoiceGenerator {
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
	invoice := types.Invoice{
		OBUId:         obuId,
		TotalDistance: totalDistance,
		Amount:        totalDistance * 10.0,
	}
	utils.WriteJsonResponse(w, http.StatusOK, invoice)
}
