package main

import "net/http"

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
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "necklace-to-bond - @milkyholme on Discord")
	resp, err := http.DefaultClient.Do(req)
}
