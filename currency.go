package mint

import (
	"errors"
	"strings"
)

type Currency uint16

// Available currencies
const (
	EUR Currency = iota
	USD
)

// FromString parses a string into a Currency (external use only).
func FromString(s string) (Currency, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	switch s {
	case "EUR":
		return EUR, nil
	case "USD":
		return USD, nil
	default:
		return 0, errors.New("invalid currency: " + s)
	}
}

// String returns the ISO currency code.
func (c Currency) String() string {
	switch c {
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	default:
		return ""
	}
}

// isValid reports whether the currency is in the enum.
func (c Currency) isValid() bool {
	return c == EUR || c == USD
}

// DecimalPlaces returns the standard number of fraction digits for a currency
func (c Currency) DecimalPlaces() (int, bool) {
	switch c {
	case EUR, USD:
		return 2, true
	default:
		return 0, false
	}
}
