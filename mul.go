package mint

import "math/big"

// Mul multiplies the amount by an integer; scale is preserved.
func (m Mint) Mul(n int64) Mint {
	factor := big.NewInt(n)
	amount := new(big.Int).Mul(m.Amount, factor)
	return Mint{Amount: amount, Currency: m.Currency, Scale: m.Scale}
}
