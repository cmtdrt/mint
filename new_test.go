package mint_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestNewFromStrValid(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"10", "10"},
		{"10.92", "10.92"},
		{"0.01", "0.01"},
		{"-10.92", "-10.92"},
		{"0.00", "0.00"},
		{"0", "0"},
		{"-0", "0"},
		{"  42.5  ", "42.5"},
		{"00010.50", "10.50"},
		{"100", "100"},
		{"0.1", "0.1"},
		{"-0.01", "-0.01"},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			m, err := mint.NewFromStr(tc.input, mint.EUR)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertString(t, m, tc.want)
			if err := m.Validate(); err != nil {
				t.Errorf("Validate: %v", err)
			}
		})
	}
}

func TestNewFromStrInvalid(t *testing.T) {
	invalid := []string{
		"10..2",
		"abc",
		"",
		"10.-2",
		".",
		"1.2.3",
		"--1",
		"+10",
		"10e2",
		" ",
		"-",
	}

	for _, input := range invalid {
		t.Run(input, func(t *testing.T) {
			_, err := mint.NewFromStr(input, mint.EUR)
			if !errors.Is(err, mint.ErrInvalidAmount) {
				t.Errorf("err = %v, want ErrInvalidAmount", err)
			}
		})
	}
}

func TestNewBigInt(t *testing.T) {
	amount := big.NewInt(1092)
	m := mint.New(amount, mint.EUR, 2)
	assertString(t, m, "10.92")

	amount.SetInt64(0)
	assertString(t, m, "10.92")
}

func TestNewBigIntNil(t *testing.T) {
	m := mint.New(nil, mint.EUR, 0)
	if !m.IsZero() {
		t.Error("nil big.Int should produce zero amount")
	}
}

func TestNewBigIntNegativeScaleValidate(t *testing.T) {
	m := mint.New(big.NewInt(100), mint.EUR, -1)
	if err := m.Validate(); err == nil {
		t.Error("expected validation error for negative scale")
	}
}
