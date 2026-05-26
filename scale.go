package mint

import "math/big"

func (m Mint) checkCurrency(other Mint) error {
	if m.Currency != other.Currency {
		return ErrCurrencyMismatch
	}
	return nil
}

func alignScales(a, b Mint) (*big.Int, *big.Int, int) {
	if a.Scale == b.Scale {
		return new(big.Int).Set(a.Amount), new(big.Int).Set(b.Amount), a.Scale
	}
	if a.Scale > b.Scale {
		factor := pow10(a.Scale - b.Scale)
		bv := new(big.Int).Mul(b.Amount, factor)
		return new(big.Int).Set(a.Amount), bv, a.Scale
	}
	factor := pow10(b.Scale - a.Scale)
	av := new(big.Int).Mul(a.Amount, factor)
	return av, new(big.Int).Set(b.Amount), b.Scale
}

func pow10(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}
