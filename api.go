package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const baseURL = "https://prices.runescape.wiki/api/v1/osrs/latest"

type APIResponse struct {
	Data map[string]Quote `json:"data"`
}

type Quote struct {
	High int `json:"high"`
	Low  int `json:"low"`
}

const (
	goldBarID     = 2357
	diamondNeckID = 1662
	diamondID     = 1601
	bondID        = 13190
)

func fetchItemQuote(id int) (Quote, error) {
	req, err := http.NewRequest("GET", baseURL+"?id="+strconv.Itoa(id), nil)
	if err != nil {
		return Quote{}, err
	}
	req.Header.Set("User-Agent", "necklace-to-bond - @milkyholme on Discord")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Quote{}, fmt.Errorf("item %d: bad status %s", id, resp.Status)
	}
	var response APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Quote{}, fmt.Errorf("item %d: decode: %w", id, err)
	}
	quote, ok := response.Data[strconv.Itoa(id)]
	if !ok {
		return Quote{}, fmt.Errorf("item %d: not in response", id)
	}
	return quote, nil
}
