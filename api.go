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

func fetchAllQuotes() (map[int]Quote, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "necklace-to-bond - @milkyholme on Discord")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status %s", resp.Status)
	}
	var response APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	quotes := make(map[int]Quote, len(response.Data))
	for id, quote := range response.Data {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("item id %q: not a number", id)
		}
		quotes[intID] = quote
	}
	return quotes, nil
}
