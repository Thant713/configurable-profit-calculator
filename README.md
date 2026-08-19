# OSRS Necklace Profit

A CLI profit calculator for F2P OSRS necklaces.

## Features

- Auto-fetch prices from OSRS Wiki API
- Display all 5 F2P necklaces sorted by ROI
- Columns: Item Name | ROI% | Profit/Item | Sell Price | Gold Bar Price | Gem Price
- 2% GE tax calculation (capped at 5M GP, 0 GP under 50 GP)

## Preset Necklaces

- Gold necklace — 1 ingredient: Gold bar
- Sapphire necklace — 2 ingredients: Gold bar + Sapphire
- Emerald necklace — 2 ingredients: Gold bar + Emerald
- Ruby necklace — 2 ingredients: Gold bar + Ruby
- Diamond necklace — 2 ingredients: Gold bar + Diamond

## Tech Stack

- **Language:** Go

## Status

Currently in development — API fetch and display in progress.

## License

MIT

## How to Run

```bash
go run .
```
