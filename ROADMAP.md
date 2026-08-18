# Configurable Profit Calculator — Roadmap

### Startup

- Prompt: "Choose mode: 1. Manual 2. Auto-fetch 1 necklace 3. Auto-fetch all necklaces"

### Phase 1: Manual Profit Calculator

- User inputs Product's Instant Buy price(user is selling at), Ingredient 1's Instant sell price(user is buying at), Ingredient 2's Instant sell price(user is buying at)
- Error and exit if sell price exceeds 250,000,000 gp
- 2% GE tax applied to the Product's Instant Buy price(user is selling at), rounded down to the nearest whole gp
- Prints results to terminal in fixed columns:
  Profit/Item | Product's Instant Buy price(user is selling at) | Ingredient 1's Instant sell price(user is buying at) | Ingredient 2's Instant sell price(user is buying at)
- Column positions are fixed — ingredient 1 and 2 never swap
- If item is a known craftable necklace (stored in necklaces.db), updates that necklace's row
- Custom items display results in terminal only; nothing saved

### Phase 2: Automatic Profit Calculator for Each Necklace

- User inputs an item name (e.g., diamond necklace)
- Go fetches current high/low prices from the OSRS Wiki Real-time Prices API
- Does phase 1 tasks automatically
- Prints connection/API error if the API is unavailable

### Phase 3: Automatic Profit Calculator for All Craftable Necklaces

- Lists all craftable necklaces defined in necklaces.db
- Does phase 2 tasks automatically
