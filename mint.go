package mint

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

type Mint struct {
	Amount   *big.Int
	Currency Currency
	Scale    int
}

func New(amount *big.Int, currency Currency, scale int) *Mint {
	return &Mint{
		Amount:   amount,
		Currency: currency,
		Scale:    scale,
	}
}

func NewFromStr(amount string, currency Currency) (*Mint, error) {
	// Allowed: optional minus sign, followed by digits and 0 or 1 dot
	re := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	if !re.MatchString(amount) {
		return nil, fmt.Errorf("invalid amount: %q", amount)
	}

	// Normalize: remove trailing zeros after the dot, and a potential lone dot after
	if strings.Contains(amount, ".") {
		amount = strings.TrimRight(amount, "0")
		amount = strings.TrimRight(amount, ".")
	}

	// If there is no dot, create a bigInt with the amount and set scale to 0
	dotIdx := strings.Index(amount, ".")
	if dotIdx == -1 {
		n := new(big.Int)
		n.SetString(amount, 10)
		return New(n, currency, 0), nil
	}

	// Since there is a dot, the scale is the length of the string after the dot
	// and the amount will be a bigInt of amount without the dot
	// example : "148.32" --> BigInt("14832") and scale 2
	scale := len(amount) - dotIdx - 1
	ammountWithoutDot := strings.Replace(amount, ".", "", 1)
	amountBigInt := new(big.Int)
	amountBigInt.SetString(ammountWithoutDot, 10)
	return New(amountBigInt, currency, scale), nil
}

func (m *Mint) IsNegative() bool {
	return m.Amount.Sign() < 0
}

func (m *Mint) IsPositive() bool {
	return m.Amount.Sign() > 0
}

func (m *Mint) IsZero() bool {
	return m.Amount.Sign() == 0
}

func (m *Mint) IsNotZero() bool {
	return m.Amount.Sign() != 0
}
