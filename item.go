package main

type Item struct {
	name string
	buy  int
	sell int
}

func item(name string, q Quote) Item {
	return Item{name: name, buy: q.Low, sell: q.High}
}

type Recipe struct {
	product     Item
	ingredients []Item
}

func (r Recipe) profit(taxAmt int) int {
	total := r.product.sell
	for _, ing := range r.ingredients {
		total -= ing.buy
	}
	return total - taxAmt
}

func tax(price int) int {
	if price < 50 {
		return 0
	}
	taxAmt := price * 2 / 100
	return min(taxAmt, 5_000_000)
}
