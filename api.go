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
	goldBarID      = 2357
	goldNeckID     = 1654
	sapphireNeckID = 1656
	emeraldNeckID  = 1658
	rubyNeckID     = 1660
	diamondNeckID  = 1662
	sapphireID     = 1607
	emeraldID      = 1605
	rubyID         = 1603
	diamondID      = 1601
)

func fetchItemPrice(id int) (p PriceEntry, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "osrs-necklace-profit - @milkyholme on Discord")
	resp, err := http.DefaultClient.Do(req)
}
