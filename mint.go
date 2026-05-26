package mint

import "math/big"

// Mint is an immutable monetary amount (scaled integer, no floats).
type Mint struct {
	Amount   *big.Int
	Currency Currency
	Scale    int
}
