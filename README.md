# OSRS Necklace Profit

A CLI tool for F2P OSRS: see diamond necklace crafting profit and how far it gets you toward a bond.

## Features

- Fetches live prices for all items in one request from the OSRS Wiki API
- Calculates profit per diamond necklace after the 2% GE tax
- Prints an aligned table: product, profit, sell price, both buy prices, bond price

## Example Output

```
Product           Profit  Sell  Buy1  Buy2  Bond
Diamond necklace  225     2014  84    1665  11365923
```

## Tech Stack

- **Language:** Go — standard library only, no dependencies

## How to Run

```
go run .
```

Needs an internet connection (fetches live prices).

## Status

v1 works end-to-end.

Next up:

## License

MIT
