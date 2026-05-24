package mint

import (
	"errors"
	"strings"
)

type Currency uint16

// Available currencies
const (
	USD Currency = iota // US Dollar
	EUR                 // Euro
)

// Convert a string to a Currency
func FromString(s string) (Currency, error) {
	s = strings.ToUpper(s)
	switch s {
	case "USD":
		return USD, nil
	case "EUR":
		return EUR, nil
	default:
		return 0, errors.New("invalid currency : " + s)
	}
}

// Convert a Currency to a string
func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	}
	return ""
}
