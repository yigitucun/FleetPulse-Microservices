package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FleetClient struct {
	BaseURL string
}

func NewFleetClient(url string) *FleetClient {
	return &FleetClient{BaseURL: url}
}

func (c *FleetClient) ValidateTruck(truckID string) (bool, error) {
	url := fmt.Sprintf("%s/api/v1/fleet/validate/%s", c.BaseURL, truckID)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	var isValid bool
	if err := json.NewDecoder(resp.Body).Decode(&isValid); err != nil {
		return false, err
	}

	return isValid, nil
}