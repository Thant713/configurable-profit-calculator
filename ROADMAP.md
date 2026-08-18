# Configurable Profit Calculator — Roadmap

### Startup

- Prompt: "Choose mode: 1) Adjust items 2) View items"

### Phase 1: Adjust Items

- 1a) Add — prompt for item name, sell price, buy price 1, buy price 2; insert into items.db
- 1b) Edit — list items, select one, update fields
- 1c) Remove — list items, select one, delete from items.db
- 2% GE tax applied to sell price, capped at 5,000,000 GP
- Tax: 0 GP if sell price < 50 GP
- Profit per item calculated based on tax

### Phase 2: View Items

- 2a) View existing items — display all items from items.db using sqlite3
- 2b) Update existing items — fetch current prices from OSRS Wiki API for all items, update items.db, display results
- Prints connection/API error if the API is unavailable
