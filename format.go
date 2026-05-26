package mint

// FormattedMint is a formatted amount with its currency, useful for display.
type FormattedMint struct {
	Amount   string
	Currency Currency
}

// Format the Mint into a formatted string with the currency
// Example : Mint(Amount:1099, Currency:EUR, Scale:2) --> FormattedMint{Amount: "10.99", Currency: EUR}
func (m Mint) Format() FormattedMint {
	raw := m.Amount.String()

	if m.Scale == 0 {
		return FormattedMint{Amount: raw, Currency: m.Currency}
	}

	// Handle negative: strip the minus, reinsert it after
	negative := m.Amount.Sign() < 0
	digits := raw
	if negative {
		digits = raw[1:]
	}

	// Pad with leading zeros if digits are fewer than scale
	// ex: scale=4, digits="3" --> "00003"
	for len(digits) <= m.Scale {
		digits = "0" + digits
	}

	dotPos := len(digits) - m.Scale
	formatted := digits[:dotPos] + "." + digits[dotPos:]
	if negative {
		formatted = "-" + formatted
	}

	return FormattedMint{Amount: formatted, Currency: m.Currency}
}

// String returns the amount part of Format() (no currency).
func (m Mint) String() string {
	return m.Format().Amount
}
