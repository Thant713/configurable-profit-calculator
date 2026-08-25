# OSRS Necklace Profit

A CLI tool for F2P OSRS: see diamond necklace crafting profit and bond price.

## Features

- Fetches live prices for all items in one request from the OSRS Wiki API
- Calculates profit per diamond necklace after the 2% GE tax
- Prints a table: product name, product sell price, ingredient 1 buy and sell prices, ingredient 2 buy and sell prices, profit per item after tax, bond buy price (buy low, sell high)
- Prices from the [OSRS Wiki real-time API](https://oldschool.runescape.wiki/w/RuneScape:Real-time_Prices)

## Example Output

```
Product   pSell  i1Buy  i1Sell  i2Buy  i2Sell  Profit  Bond
dia nklc  1995   86     89      1686   1700    184     11701284
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
