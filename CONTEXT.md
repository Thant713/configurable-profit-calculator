# OSRS Necklace Profit

A CLI tool that calculates Grand Exchange profit from crafting diamond necklaces.

## Language

**Necklace**:
A crafted item made from a gold bar and a gem. The user sells this on the Grand Exchange.
_Avoid_: product, output

**Bond**:
A membership item whose price is displayed alongside necklace profit for reference.
_Avoid_: membership token

**GE tax**:
2% of the sell price when selling items on the Grand Exchange, if the sell price is >= 50 GP. Capped at 5,000,000 GP.
_Avoid_: exchange tax, transaction fee

**High price**:
The instant sell price on the Grand Exchange — what buyers are currently paying for an item.
_Avoid_: sell price, bid price

**Low price**:
The instant buy price on the Grand Exchange — what sellers are currently asking for an item.
_Avoid_: buy price, offer price

**Profit**:
Revenue minus costs minus GE tax, calculated per necklace.
_Avoid_: margin, gain

**Recipe**:
The materials needed to craft a necklace: one gold bar plus one gem.
_Avoid_: ingredients, components
