package mint

import (
	"fmt"
	"math/big"
)

// Tax returns the tax amount for m using the given ratio rate (e.g. "0.20" for 20%).
func (m Mint) Tax(rateStr string) (Mint, error) {
	r, err := parseRate(rateStr)
	if err != nil {
		return Mint{}, err
	}
	if r.isNegative() {
		return Mint{}, fmt.Errorf("mint: %w", ErrInvalidOperation)
	}

	amount := mulAmountByRate(m.Amount, r)
	return Mint{Amount: amount, Currency: m.Currency, Scale: m.Scale + r.scale}, nil
}

// TaxExclusive returns m plus its tax (m is assumed tax-exclusive).
func (m Mint) TaxExclusive(rateStr string) (Mint, error) {
	t, err := m.Tax(rateStr)
	if err != nil {
		return Mint{}, err
	}
	return m.Add(t)
}

// TaxInclusive returns the net amount when m already includes tax.
// It computes net = gross / (1 + rate) exactly; if the decimal expansion repeats, it returns ErrNonTerminatingResult.
func (m Mint) TaxInclusive(rateStr string) (Mint, error) {
	r, err := parseRate(rateStr)
	if err != nil {
		return Mint{}, err
	}
	if r.isNegative() {
		return Mint{}, fmt.Errorf("mint: %w", ErrInvalidOperation)
	}

	den := r.addOne() // 1 + rate
	q, scale, err := divAmountByRate(m.Amount, m.Scale, den)
	if err != nil {
		return Mint{}, err
	}
	return Mint{Amount: q, Currency: m.Currency, Scale: scale}, nil
}

func mulAmountByRate(amount *big.Int, r rate) *big.Int {
	return new(big.Int).Mul(amount, r.value)
}

// divAmountByRate divides amount (with scale) by a decimal rate (value/10^scale) exactly if possible.
// Returns a new integer amount and its resulting scale. Scale may increase to represent an exact quotient.
func divAmountByRate(amount *big.Int, amountScale int, den rate) (*big.Int, int, error) {
	// quotient = (A/10^aS) / (D/10^dS) = A*10^dS / (D*10^aS)
	// We return it as an integer with initial scale = amountScale:
	// qInt = (A*10^dS) / D  with output scale = amountScale, then extend scale if remainder != 0.
	negResult := (amount.Sign() < 0) != (den.value.Sign() < 0)
	aAbs := new(big.Int).Abs(amount)
	dAbs := new(big.Int).Abs(den.value)

	numer := new(big.Int).Mul(aAbs, pow10(den.scale))
	q := new(big.Int).Quo(numer, dAbs)
	rem := new(big.Int).Rem(numer, dAbs)

	scale := amountScale
	if rem.Sign() == 0 {
		if negResult {
			q.Neg(q)
		}
		return q, scale, nil
	}

	const maxExtraDigits = 38
	ten := big.NewInt(10)
	for range maxExtraDigits {
		rem.Mul(rem, ten)
		digit := new(big.Int).Quo(rem, dAbs)
		rem.Rem(rem, dAbs)
		q.Mul(q, ten)
		q.Add(q, digit)
		scale++
		if rem.Sign() == 0 {
			if negResult {
				q.Neg(q)
			}
			return q, scale, nil
		}
	}

	return nil, 0, fmt.Errorf("mint: %w", ErrNonTerminatingResult)
}

