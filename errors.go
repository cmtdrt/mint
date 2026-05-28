package mint

import "errors"

var (
	// ErrCurrencyMismatch is returned when two amounts have different currencies.
	ErrCurrencyMismatch = errors.New("currency mismatch")
	// ErrDivisionByZero is returned when dividing by zero.
	ErrDivisionByZero = errors.New("division by zero")
	// ErrInvalidAmount is returned when the input string is not a valid amount.
	ErrInvalidAmount = errors.New("invalid amount")
	// ErrInvalidRate is returned when the input string is not a valid rate.
	ErrInvalidRate = errors.New("invalid rate")
	// ErrNonTerminatingResult is returned when an exact decimal result would be repeating.
	ErrNonTerminatingResult = errors.New("non-terminating result")
	// ErrInvalidOperation is returned when an operation's inputs are invalid (ex: negative tax rate).
	ErrInvalidOperation = errors.New("invalid operation")
)
