package mint

import "fmt"

// DiscountPercent returns m after applying a discount ratio (e.g. "0.20" for 20% off).
func (m Mint) DiscountPercent(rateStr string) (Mint, error) {
	r, err := parseRate(rateStr)
	if err != nil {
		return Mint{}, err
	}
	if r.isNegative() {
		return Mint{}, fmt.Errorf("mint: %w", ErrInvalidOperation)
	}
	if r.cmpOne() > 0 {
		return Mint{}, fmt.Errorf("mint: %w", ErrInvalidOperation)
	}

	disc, err := m.Tax(rateStr) // same math: amount * rate
	if err != nil {
		return Mint{}, err
	}
	return m.Sub(disc)
}

// DiscountAmount returns m after subtracting a fixed discount amount (same currency required).
func (m Mint) DiscountAmount(amount Mint) (Mint, error) {
	if err := m.checkCurrency(amount); err != nil {
		return Mint{}, err
	}
	return m.Sub(amount)
}

