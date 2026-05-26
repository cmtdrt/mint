package test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestNewBigInt(t *testing.T) {
	amount := big.NewInt(1092)
	m := mint.New(amount, mint.EUR, 2)
	if m.String() != "10.92" {
		t.Errorf("New(big.Int).String() = %q, want 10.92", m.String())
	}
	// Internal copy must not change the Mint if the caller mutates their big.Int.
	amount.SetInt64(0)
	if m.String() != "10.92" {
		t.Errorf("after mutating source big.Int: got %q", m.String())
	}
}

func TestNewValid(t *testing.T) {
	cases := []struct {
		input    string
		currency mint.Currency
		want     string
	}{
		{"10", mint.EUR, "10"},
		{"10.92", mint.EUR, "10.92"},
		{"0.01", mint.USD, "0.01"},
		{"-10.92", mint.EUR, "-10.92"},
		{"0.00", mint.EUR, "0.00"},
	}

	for _, tc := range cases {
		m, err := mint.NewFromStr(tc.input, tc.currency)
		if err != nil {
			t.Fatalf("New(%q): %v", tc.input, err)
		}
		if got := m.String(); got != tc.want {
			t.Errorf("New(%q).String() = %q, want %q", tc.input, got, tc.want)
		}
		if err := m.Validate(); err != nil {
			t.Errorf("Validate(%q): %v", tc.input, err)
		}
	}
}

func TestNewInvalid(t *testing.T) {
	invalid := []string{"10..2", "abc", "", "10.-2", "."}
	for _, input := range invalid {
		_, err := mint.NewFromStr(input, mint.EUR)
		if !errors.Is(err, mint.ErrInvalidAmount) {
			t.Errorf("New(%q) err = %v, want ErrInvalidAmount", input, err)
		}
	}
}

func TestAddSub(t *testing.T) {
	a, _ := mint.NewFromStr("10.50", mint.EUR)
	b, _ := mint.NewFromStr("0.5", mint.EUR)

	sum, err := a.Add(b)
	if err != nil {
		t.Fatal(err)
	}
	if sum.String() != "11.00" {
		t.Errorf("Add: got %q, want 11.00", sum.String())
	}

	c, _ := mint.NewFromStr("3.00", mint.USD)
	diff, err := sum.Sub(c)
	if err == nil {
		t.Fatal("expected currency mismatch")
	}
	if !errors.Is(err, mint.ErrCurrencyMismatch) {
		t.Errorf("Sub err = %v", err)
	}

	_ = diff
}

func TestMulDiv(t *testing.T) {
	m, _ := mint.NewFromStr("10.92", mint.EUR)
	triple := m.Mul(3)
	if triple.String() != "32.76" {
		t.Errorf("Mul: got %q", triple.String())
	}

	half, err := m.Div(2)
	if err != nil {
		t.Fatal(err)
	}
	if half.String() != "5.46" {
		t.Errorf("Div: got %q", half.String())
	}

	exact, err := mint.NewFromStr("10.00", mint.EUR)
	if err != nil {
		t.Fatal(err)
	}
	third, err := exact.Div(4)
	if err != nil {
		t.Fatal(err)
	}
	if third.String() != "2.50" {
		t.Errorf("Div exact: got %q", third.String())
	}
}

func TestEqualIsZero(t *testing.T) {
	a, _ := mint.NewFromStr("1.0", mint.EUR)
	b, _ := mint.NewFromStr("1.00", mint.EUR)
	if !a.Equal(b) {
		t.Error("1.0 should equal 1.00")
	}

	z, _ := mint.NewFromStr("0", mint.EUR)
	if !z.IsZero() {
		t.Error("zero should be zero")
	}
}

func TestDivByZero(t *testing.T) {
	m, _ := mint.NewFromStr("1", mint.EUR)
	_, err := m.Div(0)
	if !errors.Is(err, mint.ErrDivisionByZero) {
		t.Errorf("Div(0) err = %v", err)
	}
}
