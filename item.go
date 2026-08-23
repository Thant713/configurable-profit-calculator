package main

type Item struct {
	name      string
	sellPrice int
	gemPrice  int
	goldPrice int
}

func profit(it Item, tax int) int {
	return it.sellPrice - it.gemPrice - it.goldPrice - tax
}

func tax(price int) int {
	if price < 50 {
		return 0
	}
	taxAmt := price * 2 / 100
	return min(taxAmt, 5_000_000)
}
