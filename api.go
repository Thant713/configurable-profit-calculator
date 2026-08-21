package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const baseURL = "https://prices.runescape.wiki/api/v1/osrs/latest"

type APIResponse struct {
	Data map[string]PriceEntry `json:"data"`
}

type PriceEntry struct {
	High int `json:"high"`
	Low  int `json:"low"`
}

const (
	goldBarID     = 2357
	diamondNeckID = 1662
	diamondID     = 1601
)

func fetchItemPrice(id int) (p PriceEntry, err error) {
	req, err := http.NewRequest("GET", baseURL+"?id="+strconv.Itoa(id), nil)
	if err != nil {
		return PriceEntry{}, err
	}
	req.Header.Set("User-Agent", "necklace-to-bond - @milkyholme on Discord")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return PriceEntry{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return PriceEntry{}, fmt.Errorf("item %d: bad status %s", id, resp.Status)
	}
	var r APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return PriceEntry{}, fmt.Errorf("item %d: decode: %w", id, err)
	}
	p, ok := r.Data[strconv.Itoa(id)]
	if !ok {
		return PriceEntry{}, fmt.Errorf("item %d: not in response", id)
	}
	return p, nil
}
