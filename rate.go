package mint

import (
	"math/big"
	"strings"
)

// rate is a currency-less decimal ratio (ex: 0.20 for 20%).
type rate struct {
	value *big.Int
	scale int
}

func parseRate(input string) (rate, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return rate{}, ErrInvalidRate
	}

	negative := false
	if strings.HasPrefix(input, "-") {
		negative = true
		input = input[1:]
		if input == "" {
			return rate{}, ErrInvalidRate
		}
	}

	parts := strings.Split(input, ".")
	if len(parts) > 2 {
		return rate{}, ErrInvalidRate
	}

	intPart := parts[0]
	var decPart string
	scale := 0

	if len(parts) == 2 {
		decPart = parts[1]
		scale = len(decPart)
		for _, r := range decPart {
			if r < '0' || r > '9' {
				return rate{}, ErrInvalidRate
			}
		}
	}

	if intPart == "" {
		return rate{}, ErrInvalidRate
	}
	for _, r := range intPart {
		if r < '0' || r > '9' {
			return rate{}, ErrInvalidRate
		}
	}

	intPart = trimLeadingZeros(intPart)
	combined := intPart + decPart
	if combined == "" {
		combined = "0"
	}

	v := new(big.Int)
	if _, ok := v.SetString(combined, 10); !ok {
		return rate{}, ErrInvalidRate
	}
	if negative {
		v.Neg(v)
	}

	return rate{value: v, scale: scale}, nil
}

func (r rate) isNegative() bool {
	return r.value.Sign() < 0
}

func (r rate) isZero() bool {
	return r.value.Sign() == 0
}

// cmpOne compares r to 1.0.
func (r rate) cmpOne() int {
	one := pow10(r.scale) // 1.0 in scaled form
	return r.value.Cmp(one)
}

func (r rate) addOne() rate {
	one := pow10(r.scale)
	return rate{value: new(big.Int).Add(r.value, one), scale: r.scale}
}
