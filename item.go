package main

type Item struct {
	name  string
	price int
}

type Recipe struct {
	product     Item
	ingredients []Item
}

func (r Recipe) profit(taxAmt int) int {
	total := r.product.price
	for _, ing := range r.ingredients {
		total -= ing.price
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
