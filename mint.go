package mint

import (
	"fmt"
	"math/big"
	"strings"
)

// Mint is an immutable monetary amount (scaled integer, no floats).
type Mint struct {
	Amount   *big.Int
	Currency Currency
	Scale    int
}

// New builds a Mint from a scaled integer (copies amount).
func New(amount *big.Int, currency Currency, scale int) Mint {
	amt := new(big.Int)
	if amount != nil {
		amt.Set(amount)
	}
	return Mint{Amount: amt, Currency: currency, Scale: scale}
}

// NewFromStr builds a Mint from a decimal string and a currency.
func NewFromStr(input string, currency Currency) (Mint, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return Mint{}, ErrInvalidAmount
	}

	negative := false
	if strings.HasPrefix(input, "-") {
		negative = true
		input = input[1:]
		if input == "" {
			return Mint{}, ErrInvalidAmount
		}
	}

	parts := strings.Split(input, ".")
	if len(parts) > 2 {
		return Mint{}, ErrInvalidAmount
	}

	intPart := parts[0]
	var decPart string
	scale := 0

	if len(parts) == 2 {
		decPart = parts[1]
		scale = len(decPart)
		for _, r := range decPart {
			if r < '0' || r > '9' {
				return Mint{}, ErrInvalidAmount
			}
		}
	}

	if intPart == "" {
		return Mint{}, ErrInvalidAmount
	}

	for _, r := range intPart {
		if r < '0' || r > '9' {
			return Mint{}, ErrInvalidAmount
		}
	}

	intPart = trimLeadingZeros(intPart)
	combined := intPart + decPart
	if combined == "" {
		combined = "0"
	}

	amount := new(big.Int)
	if _, ok := amount.SetString(combined, 10); !ok {
		return Mint{}, ErrInvalidAmount
	}
	if negative {
		amount.Neg(amount)
	}

	return Mint{Amount: amount, Currency: currency, Scale: scale}, nil
}

func trimLeadingZeros(s string) string {
	if s == "" || s == "0" {
		return "0"
	}
	i := 0
	for i < len(s)-1 && s[i] == '0' {
		i++
	}
	return s[i:]
}

func (m Mint) IsNegative() bool {
	return m.Amount.Sign() < 0
}

func (m Mint) IsPositive() bool {
	return m.Amount.Sign() > 0
}

func (m Mint) IsZero() bool {
	return m.Amount.Sign() == 0
}

func (m Mint) IsNotZero() bool {
	return m.Amount.Sign() != 0
}

// Equal compares two amounts (currency and value at the same scale).
func (m Mint) Equal(other Mint) bool {
	if m.Currency != other.Currency {
		return false
	}
	a, b, _ := alignScales(m, other)
	return a.Cmp(b) == 0
}

// String returns the amount part of Format() (no currency).
func (m Mint) String() string {
	return m.Format().Amount
}

// Validate checks internal integrity of the amount.
func (m Mint) Validate() error {
	if m.Amount == nil {
		return fmt.Errorf("mint: nil amount")
	}
	if m.Scale < 0 {
		return fmt.Errorf("mint: negative scale")
	}
	if !m.Currency.isValid() {
		return fmt.Errorf("mint: invalid currency")
	}
	return nil
}

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
