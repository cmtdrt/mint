package mint

import "math/big"

// Add adds two amounts of the same currency (scale alignment, no rounding).
func (m Mint) Add(other Mint) (Mint, error) {
	if err := m.checkCurrency(other); err != nil {
		return Mint{}, err
	}
	a, b, scale := alignScales(m, other)
	sum := new(big.Int).Add(a, b)
	return Mint{Amount: sum, Currency: m.Currency, Scale: scale}, nil
}
