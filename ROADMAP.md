# Roadmap

### Data Sources

- OSRS Wiki API: `https://prices.runescape.wiki/api/v1/osrs/latest`
- Items: diamond necklace (1662), gold bar (2357), diamond (1601), bond (13190)

### Display

- Item name
- Profit per item
- Product's sell price — what the user sells it for (high field)
- Gold bar's buy price — what the user pays for it (low field)
- Gem's buy price — what the user pays for it (low field)
- Bond price (low field)

### Tax Rules

- 2% of sell price if >= 50 GP, else 0
- Tax capped at 5,000,000 GP
- Tax applies only to the necklace sell (inputs are bought, not sold)

### Profit Formula

profit = sell price - (gold bar buy + gem buy) - GE tax

## Future Considerations
