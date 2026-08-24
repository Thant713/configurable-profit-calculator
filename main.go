package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	quotes, err := fetchAllQuotes()
	if err != nil {
		log.Fatal(err)
	}

	necklace := Item{name: "Diamond necklace", price: quotes[diamondNeckID].High}
	gem := Item{name: "Diamond", price: quotes[diamondID].Low}
	bar := Item{name: "Gold bar", price: quotes[goldBarID].Low}
	bondItem := Item{name: "Bond", price: quotes[bondID].Low}

	recipe := Recipe{product: necklace, ingredients: []Item{gem, bar}}

	// code below does what I wanted cat sqlite3 to do but without needing sqlite3

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "Product\tProfit\tSell\tBuy1\tBuy2\tBond")
	p := recipe.profit(tax(necklace.price))
	fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\t%d\n", necklace.name, p, necklace.price, bar.price, gem.price, bondItem.price)
	w.Flush()
}
