package test

import (
	"testing"

	"github.com/cmtdrt/mint"
)

func TestEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want bool
	}{
		{"same representation", "10.50", "10.50", true},
		{"different scale same value", "1.0", "1.00", true},
		{"different value", "1.00", "1.01", false},
		{"zero variants", "0", "0.00", true},
		{"negative", "-2.50", "-2.50", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := mustParse(t, tc.a, mint.EUR)
			b := mustParse(t, tc.b, mint.EUR)
			if got := a.Equal(b); got != tc.want {
				t.Errorf("Equal = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestEqualCurrencyMismatch(t *testing.T) {
	a := mustParse(t, "1.00", mint.EUR)
	b := mustParse(t, "1.00", mint.USD)
	if a.Equal(b) {
		t.Error("different currencies should not be equal")
	}
}

func TestIsZero(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"0", true},
		{"0.00", true},
		{"0.01", false},
		{"-0.00", true},
		{"1", false},
	}

	for _, tc := range cases {
		m := mustParse(t, tc.in, mint.EUR)
		if got := m.IsZero(); got != tc.want {
			t.Errorf("IsZero(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestIsPositiveNegative(t *testing.T) {
	pos := mustParse(t, "0.01", mint.EUR)
	if !pos.IsPositive() || pos.IsNegative() || pos.IsZero() {
		t.Error("0.01 should be positive only")
	}

	neg := mustParse(t, "-1", mint.EUR)
	if !neg.IsNegative() || neg.IsPositive() {
		t.Error("-1 should be negative")
	}

	zero := mustParse(t, "0", mint.EUR)
	if zero.IsPositive() || zero.IsNegative() || !zero.IsZero() {
		t.Error("zero should be zero only")
	}
	if zero.IsNotZero() {
		t.Error("zero should not be IsNotZero")
	}
}

func TestIsNotZero(t *testing.T) {
	if !mustParse(t, "1", mint.EUR).IsNotZero() {
		t.Error("1 should be not zero")
	}
}
