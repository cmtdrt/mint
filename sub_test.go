package mint_test

import (
	"errors"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want string
	}{
		{"same scale", "5.00", "2.00", "3.00"},
		{"different scale", "10.50", "0.5", "10.00"},
		{"to zero", "7.25", "7.25", "0.00"},
		{"negative result", "1.00", "3.00", "-2.00"},
		{"minus negative", "5.00", "-2.00", "7.00"},
		{"subtract zero", "4.20", "0", "4.20"},
		{"zero minus amount", "0", "1.50", "-1.50"},
		{"small decimals", "0.03", "0.01", "0.02"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := mustParse(t, tc.a, mint.EUR)
			b := mustParse(t, tc.b, mint.EUR)
			got, err := a.Sub(b)
			if err != nil {
				t.Fatalf("Sub: %v", err)
			}
			assertString(t, got, tc.want)
		})
	}
}

func TestSubCurrencyMismatch(t *testing.T) {
	a := mustParse(t, "10.00", mint.EUR)
	b := mustParse(t, "1.00", mint.USD)
	_, err := a.Sub(b)
	if !errors.Is(err, mint.ErrCurrencyMismatch) {
		t.Errorf("err = %v, want ErrCurrencyMismatch", err)
	}
}

func TestSubImmutability(t *testing.T) {
	a := mustParse(t, "5.00", mint.EUR)
	b := mustParse(t, "1.00", mint.EUR)
	beforeA := a.String()
	beforeB := b.String()
	_, _ = a.Sub(b)
	if a.String() != beforeA || b.String() != beforeB {
		t.Error("Sub should not mutate operands")
	}
}
