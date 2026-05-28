package mint_test

import (
	"errors"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name    string
		a, b    string
		want    string
		wantErr bool
	}{
		{"same scale", "1.00", "2.00", "3.00", false},
		{"different scale", "10.50", "0.5", "11.00", false},
		{"integers", "100", "50", "150", false},
		{"zero left", "0", "5.25", "5.25", false},
		{"zero right", "3.10", "0", "3.10", false},
		{"both zero", "0.00", "0", "0.00", false},
		{"negative plus positive", "-5.00", "10.00", "5.00", false},
		{"negative plus negative", "-3.50", "-1.50", "-5.00", false},
		{"small decimals", "0.01", "0.02", "0.03", false},
		{"large values", "999999.99", "0.01", "1000000.00", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := mustParse(t, tc.a, mint.EUR)
			b := mustParse(t, tc.b, mint.EUR)
			got, err := a.Add(b)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Add: %v", err)
			}
			assertString(t, got, tc.want)
		})
	}
}

func TestAddCurrencyMismatch(t *testing.T) {
	a := mustParse(t, "10.00", mint.EUR)
	b := mustParse(t, "10.00", mint.USD)
	_, err := a.Add(b)
	if !errors.Is(err, mint.ErrCurrencyMismatch) {
		t.Errorf("err = %v, want ErrCurrencyMismatch", err)
	}
}

func TestAddImmutability(t *testing.T) {
	a := mustParse(t, "1.00", mint.EUR)
	b := mustParse(t, "2.00", mint.EUR)
	beforeA := a.String()
	beforeB := b.String()
	_, _ = a.Add(b)
	if a.String() != beforeA || b.String() != beforeB {
		t.Error("Add should not mutate operands")
	}
}
