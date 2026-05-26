package mint

import "math/big"

// Sub subtracts other from m (same currency, scale alignment, no rounding).
func (m Mint) Sub(other Mint) (Mint, error) {
	if err := m.checkCurrency(other); err != nil {
		return Mint{}, err
	}
	a, b, scale := alignScales(m, other)
	diff := new(big.Int).Sub(a, b)
	return Mint{Amount: diff, Currency: m.Currency, Scale: scale}, nil
}
