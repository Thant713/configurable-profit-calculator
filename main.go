package main

import (
	"fmt"
	"log"
)

func main() {
	items := []int{goldBarID, diamondNeckID, diamondID, bondID}
	for _, item := range items {
		bothIds, err := fetchItemPrice(item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(item, bothIds)
	}
}
