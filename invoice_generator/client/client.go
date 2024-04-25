package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/handlers"
)

type InvoiceGeneratorClient struct {
	AggregatorEndpoint string
}

func NewInvoiceGeneratorClient(endpoint string) *InvoiceGeneratorClient {
	return &InvoiceGeneratorClient{
		AggregatorEndpoint: endpoint,
	}
}

func (igc *InvoiceGeneratorClient) PostDistanceData(distance *handlers.Distance) error {
	client := &http.Client{}

	distanceBytes, err := json.Marshal(distance)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", igc.AggregatorEndpoint, bytes.NewReader(distanceBytes))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to add distance data")
	}
	return nil
}
