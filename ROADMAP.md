# Configurable Profit Calculator — Roadmap

### Preset Necklaces

- Gold necklace — 1 ingredient: Gold bar
- Sapphire necklace — 2 ingredients: Gold bar + Sapphire
- Emerald necklace — 2 ingredients: Gold bar + Emerald
- Ruby necklace — 2 ingredients: Gold bar + Ruby
- Diamond necklace — 2 ingredients: Gold bar + Diamond

### Display

- Auto-fetch prices from OSRS Wiki API on launch
- Display all 5 necklaces sorted by ROI percentage
- Columns: Item Name | ROI% | Profit/Item | Sell Price | Gold Bar Price | Gem Price
- fmt.Printf in Go for display (no sqlite3 CLI)

### Tax Rules

- 2% of sell price if >= 50 GP, else 0
- Tax capped at 5,000,000 GP

### TODO

- Fetch specific prices from OSRS Wiki API:
  - Item Name
  - ROI percentage
  - Profit/Item
  - Input Product's Instant Buy price (user is selling at)
  - Input Gold bar's Instant sell price (user is buying at)
  - Input Gem's Instant sell price (user is buying at)

### Future Considerations

- SQLite database for persistence
- Manual input mode
- Budget input
