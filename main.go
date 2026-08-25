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

	necklace := item("dia nklc", quotes[diamondNeckID])
	bar := item("Gold bar", quotes[goldBarID])
	gem := item("Diamond", quotes[diamondID])
	bondItem := item("Bond", quotes[bondID])

	recipe := Recipe{product: necklace, ingredients: []Item{gem, bar}}

	// code below does what I wanted cat sqlite3 to do but without needing sqlite3

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "Product\tpSell\ti1Buy\ti1Sell\ti2Buy\ti2Sell\tProfit\tBond")
	p := recipe.profit(tax(necklace.sell))
	fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n", necklace.name, necklace.sell, bar.buy, bar.sell, gem.buy, gem.sell, p, bondItem.buy)
	w.Flush()
}
