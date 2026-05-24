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

type FormattedMint struct {
	Amount   string
	Currency Currency
}

// Format the Mint into a formatted string with the currency
// Example : Mint(1099, EUR, 2) --> FormattedMint{Amount: "10.99", Currency: EUR}
func (m *Mint) Format() FormattedMint {
	raw := m.Amount.String() // ex: "14832" ou "-14832"

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

// TODO : handle NewFromStr case with 0.X --> should be Bigint of 1 and scale 1 ?
