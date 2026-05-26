package mint

import "errors"

var (
	// ErrCurrencyMismatch is returned when two amounts have different currencies.
	ErrCurrencyMismatch = errors.New("currency mismatch")
	// ErrDivisionByZero is returned when dividing by zero.
	ErrDivisionByZero = errors.New("division by zero")
	// ErrInvalidAmount is returned when the input string is not a valid amount.
	ErrInvalidAmount = errors.New("invalid amount")
)
