package mint

import (
	"fmt"
	"math/big"
)

// Div divides the amount by an integer without silent rounding.
// Scale may increase for an exact result; otherwise an error is returned after a limit.
func (m Mint) Div(n int64) (Mint, error) {
	if n == 0 {
		return Mint{}, ErrDivisionByZero
	}

	negResult := (m.Amount.Sign() < 0) != (n < 0)
	v := new(big.Int).Abs(m.Amount)
	dAbs := new(big.Int).SetInt64(n)
	dAbs.Abs(dAbs)

	quotient := new(big.Int).Quo(v, dAbs)
	remainder := new(big.Int).Rem(v, dAbs)
	scale := m.Scale

	if remainder.Sign() == 0 {
		if negResult {
			quotient.Neg(quotient)
		}
		return Mint{Amount: quotient, Currency: m.Currency, Scale: scale}, nil
	}

	const maxExtraDigits = 38
	ten := big.NewInt(10)
	for range maxExtraDigits {
		remainder.Mul(remainder, ten)
		digit := new(big.Int).Quo(remainder, dAbs)
		remainder.Rem(remainder, dAbs)
		quotient.Mul(quotient, ten)
		quotient.Add(quotient, digit)
		scale++
		if remainder.Sign() == 0 {
			if negResult {
				quotient.Neg(quotient)
			}
			return Mint{Amount: quotient, Currency: m.Currency, Scale: scale}, nil
		}
	}

	return Mint{}, fmt.Errorf("mint: non-terminating division")
}
