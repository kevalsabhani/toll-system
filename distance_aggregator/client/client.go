package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kevalsabhani/toll-calculator/distance_aggregator/handlers"
)

type DistanceAggregatorClient struct {
	AggregatorEndpoint string
}

func NewDistanceAggregatorClient(endpoint string) *DistanceAggregatorClient {
	return &DistanceAggregatorClient{
		AggregatorEndpoint: endpoint,
	}
}

func (dac *DistanceAggregatorClient) PostDistanceData(distance *handlers.Distance) error {
	client := &http.Client{}

	distanceBytes, err := json.Marshal(distance)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", dac.AggregatorEndpoint, bytes.NewReader(distanceBytes))
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
