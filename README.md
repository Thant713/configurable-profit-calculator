# OSRS Necklace Profit

A CLI tool for F2P OSRS: see diamond necklace crafting profit and bond price.

## Features

- Fetches live prices for all items in one request from the OSRS Wiki API
- Calculates profit per diamond necklace after the 2% GE tax
- Prints an aligned table: product, profit, sell price, both buy prices, bond price
- Prices from the [OSRS Wiki real-time API](https://oldschool.runescape.wiki/w/RuneScape:Real-time_Prices)

## Example Output

```
Product           Profit  Sell  Buy1  Buy2  Bond
Diamond necklace  225     2014  84    1665  11365923
```

## How to Run

```
go run .
```

Or install it as a command runnable from any directory:

```
go build -o ~/.local/bin/necklace-to-bond .
```

Needs an internet connection (fetches live prices).

## Web Version (for personal use)

https://thant713.github.io/necklace-to-bond/web/

## License

MIT
