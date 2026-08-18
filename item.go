package main

import "fmt"

type Item struct {
	name      string
	sellPrice int
	buyPrice1 int
	buyPrice2 int
}

func (i *Item) AddItem() {
	i.name = getStringInput("Item name: ")
	i.sellPrice = getInput("Input Product's Instant Buy price(user is selling at): ")
	i.buyPrice1 = getInput("Ingredient 1's Instant sell price(user is buying at): ")
	i.buyPrice2 = getInput("Ingredient 2's Instant sell price(user is buying at): ")
}

func getStringInput(prompt string) string {
	var val string
	fmt.Println(prompt)
	fmt.Scan(&val)
	return val
}

func getInput(prompt string) int {
	var val int
	fmt.Println(prompt)
	fmt.Scan(&val)
	return val
}

/* Add to sql
func (i *Item) ProfitPerItem() int {
	tax := i.sellPrice * 2 / 100
	if tax > 5_000_000 {
		tax = 5_000_000
	}
	if i.sellPrice >= 50 {
		return i.sellPrice - tax - i.buyPrice1 - i.buyPrice2
	}
	return i.sellPrice - i.buyPrice1 - i.buyPrice2
}
*/
