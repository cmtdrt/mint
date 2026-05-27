package mint

// RoundMode selects how fractional digits are rounded away.
type RoundMode int

const (
	// RoundHalfUp rounds 0.5 away from zero (e.g. 1.235 -> 1.24).
	RoundHalfUp RoundMode = iota
	// RoundHalfDown rounds 0.5 toward zero (e.g. 1.235 -> 1.23).
	RoundHalfDown
	// RoundHalfEven rounds 0.5 to the nearest even digit (banker's rounding).
	RoundHalfEven
	// RoundFloor rounds toward negative infinity.
	RoundFloor
	// RoundCeil rounds toward positive infinity.
	RoundCeil
)
