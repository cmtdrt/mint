package mint

// Equal compares two amounts (currency and value at the same scale).
func (m Mint) Equal(other Mint) bool {
	if m.Currency != other.Currency {
		return false
	}
	a, b, _ := alignScales(m, other)
	return a.Cmp(b) == 0
}

func (m Mint) IsNegative() bool {
	return m.Amount.Sign() < 0
}

func (m Mint) IsPositive() bool {
	return m.Amount.Sign() > 0
}

// IsZero reports whether the amount is zero.
func (m Mint) IsZero() bool {
	return m.Amount.Sign() == 0
}

func (m Mint) IsNotZero() bool {
	return m.Amount.Sign() != 0
}
