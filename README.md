## Mint

`mint` is a ![Go](https://img.shields.io/badge/-%2300ADD8.svg?logo=go&logoColor=white) library to work with monetary amounts **without ever using floating point numbers**.

### Why no floating numbers?

- Floats cannot represent many decimal values exactly (e.g. 0.1), which can produce surprising results.
- Money needs **deterministic, auditable** behavior.
- `mint` uses scaled integers (`big.Int`) to keep computations predictable.

### Features

- Money type backed by `big.Int` + scale (no floats)
- Immutable operations (`Add`, `Sub`, `Mul`, `Div`)
- Currency-safe arithmetic (EUR/USD in the current layers)
- Explicit rounding only (`Round`, `RoundToCurrency`) with multiple rounding modes
- Tax and discount helpers using a ratio rate string (e.g. `"0.20"` = 20%)
- Deterministic `Split` and `Allocate` with explicit rounding mode, preserving totals