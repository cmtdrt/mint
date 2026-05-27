package mint_test

import (
	"testing"

	"github.com/cmtdrt/mint"
)

func TestCurrencyDecimalPlaces(t *testing.T) {
	if p, ok := mint.EUR.DecimalPlaces(); !ok || p != 2 {
		t.Errorf("EUR DecimalPlaces = %d, %v", p, ok)
	}
	if p, ok := mint.USD.DecimalPlaces(); !ok || p != 2 {
		t.Errorf("USD DecimalPlaces = %d, %v", p, ok)
	}
}

func TestRoundHalfUp(t *testing.T) {
	cases := []struct {
		in, want string
		scale    int
	}{
		{"10.929", "10.93", 2},
		{"10.921", "10.92", 2},
		{"10.925", "10.93", 2},
		{"-10.925", "-10.93", 2},
		{"0.005", "0.01", 2},
		{"0.004", "0.00", 2},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			m := mustParse(t, tc.in, mint.EUR)
			got := m.Round(tc.scale, mint.RoundHalfUp)
			assertString(t, got, tc.want)
		})
	}
}

func TestRoundHalfDown(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"10.925", "10.92"},
		{"10.929", "10.93"},
		{"-10.925", "-10.92"},
	}

	for _, tc := range cases {
		m := mustParse(t, tc.in, mint.EUR)
		got := m.Round(2, mint.RoundHalfDown)
		assertString(t, got, tc.want)
	}
}

func TestRoundHalfEven(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"10.125", "10.12"},
		{"10.135", "10.14"},
		{"10.925", "10.92"},
		{"10.935", "10.94"},
	}

	for _, tc := range cases {
		m := mustParse(t, tc.in, mint.EUR)
		got := m.Round(2, mint.RoundHalfEven)
		assertString(t, got, tc.want)
	}
}

func TestRoundFloorCeil(t *testing.T) {
	cases := []struct {
		in       string
		wantFloor string
		wantCeil  string
	}{
		{"10.929", "10.92", "10.93"},
		{"-10.929", "-10.93", "-10.92"},
		{"10.920", "10.92", "10.92"},
	}

	for _, tc := range cases {
		m := mustParse(t, tc.in, mint.EUR)
		assertString(t, m.Round(2, mint.RoundFloor), tc.wantFloor)
		assertString(t, m.Round(2, mint.RoundCeil), tc.wantCeil)
	}
}

func TestRoundScaleIncrease(t *testing.T) {
	m := mustParse(t, "10.5", mint.EUR)
	got := m.Round(3, mint.RoundHalfUp)
	assertString(t, got, "10.500")
	if got.Scale != 3 {
		t.Errorf("Scale = %d, want 3", got.Scale)
	}
}

func TestRoundSameScaleNoOp(t *testing.T) {
	m := mustParse(t, "10.92", mint.EUR)
	got := m.Round(2, mint.RoundHalfUp)
	assertString(t, got, "10.92")
}

func TestRoundToCurrency(t *testing.T) {
	m := mustParse(t, "10.926", mint.EUR)
	got := m.RoundToCurrency(mint.RoundHalfUp)
	assertString(t, got, "10.93")
	if got.Scale != 2 {
		t.Errorf("Scale = %d, want 2", got.Scale)
	}
}

func TestRoundToCurrencyUSD(t *testing.T) {
	m := mustParse(t, "99.999", mint.USD)
	got := m.RoundToCurrency(mint.RoundHalfUp)
	assertString(t, got, "100.00")
}

func TestRoundImmutability(t *testing.T) {
	m := mustParse(t, "10.999", mint.EUR)
	before := m.String()
	_ = m.RoundToCurrency(mint.RoundHalfUp)
	if m.String() != before {
		t.Error("Round must not mutate receiver")
	}
}

func TestRoundNegativeScale(t *testing.T) {
	m := mustParse(t, "10.92", mint.EUR)
	got := m.Round(-1, mint.RoundHalfUp)
	assertString(t, got, "10.92")
}

func TestArithmeticDoesNotRound(t *testing.T) {
	a := mustParse(t, "10.50", mint.EUR)
	b := mustParse(t, "0.5", mint.EUR)
	sum, err := a.Add(b)
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, sum, "11.00")

	m := mustParse(t, "10.00", mint.EUR)
	_, err = m.Div(3)
	if err == nil {
		t.Fatal("Div should not round; non-terminating division expected")
	}
}
