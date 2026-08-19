# OSRS Necklace Profit — Roadmap

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
- fmt.Printf in Go for display

### Tax Rules

- 2% of sell price if >= 50 GP, else 0
- Tax capped at 5,000,000 GP

### Column Definitions

- Item name
- ROI percentage
- Profit per item
- Product's instant buy price (user is selling at)
- Gold bar's instant sell price (user is buying at)
- Gem's instant sell price (user is buying at)

### Future Considerations

- SQLite database for persistence
- Manual inputs and expanding what you can calculate within the cli
