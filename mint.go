package mint

import (
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
