package mint

import "math/big"

// Round rounds m to targetScale using mode. Rounding is explicit; scale never decreases without a call to Round.
// If targetScale is greater than m.Scale, trailing zeros are added (no rounding).
func (m Mint) Round(targetScale int, mode RoundMode) Mint {
	if targetScale < 0 {
		return New(m.Amount, m.Currency, m.Scale)
	}
	if targetScale >= m.Scale {
		if targetScale == m.Scale {
			return New(m.Amount, m.Currency, m.Scale)
		}
		diff := targetScale - m.Scale
		factor := pow10(diff)
		amount := new(big.Int).Mul(m.Amount, factor)
		return Mint{Amount: amount, Currency: m.Currency, Scale: targetScale}
	}

	amount := roundAmount(m.Amount, m.Scale, targetScale, mode)
	return Mint{Amount: amount, Currency: m.Currency, Scale: targetScale}
}

// RoundToCurrency rounds m to the standard fraction digits of its currency (EUR/USD: 2).
func (m Mint) RoundToCurrency(mode RoundMode) Mint {
	places, ok := m.Currency.DecimalPlaces()
	if !ok {
		return New(m.Amount, m.Currency, m.Scale)
	}
	return m.Round(places, mode)
}

func roundAmount(amount *big.Int, scale, targetScale int, mode RoundMode) *big.Int {
	negative := amount.Sign() < 0
	abs := new(big.Int).Abs(amount)

	drop := scale - targetScale
	divisor := pow10(drop)
	q := new(big.Int).Quo(abs, divisor)
	r := new(big.Int).Rem(abs, divisor)

	if r.Sign() == 0 {
		if negative {
			q.Neg(q)
		}
		return q
	}

	if shouldIncrement(q, r, divisor, mode, negative) {
		q.Add(q, big.NewInt(1))
	}
	if negative {
		q.Neg(q)
	}
	return q
}

func shouldIncrement(q, r, divisor *big.Int, mode RoundMode, negative bool) bool {
	switch mode {
	case RoundFloor:
		if negative {
			return true
		}
		return false
	case RoundCeil:
		if negative {
			return false
		}
		return true
	case RoundHalfDown:
		return compareDoubleR(r, divisor) > 0
	case RoundHalfEven:
		cmp := compareDoubleR(r, divisor)
		if cmp > 0 {
			return true
		}
		if cmp < 0 {
			return false
		}
		return q.Bit(0) == 1
	default: // RoundHalfUp
		return compareDoubleR(r, divisor) >= 0
	}
}

// compareDoubleR compares 2*r with divisor (-1, 0, 1).
func compareDoubleR(r, divisor *big.Int) int {
	double := new(big.Int).Lsh(r, 1)
	return double.Cmp(divisor)
}
