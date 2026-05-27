package mint_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestDiv(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		n       int64
		want    string
		wantErr string
	}{
		{"exact half", "10.92", 2, "5.46", ""},
		{"exact quarter", "10.00", 4, "2.50", ""},
		{"integer division", "100", 4, "25", ""},
		{"scale increases terminating", "1.00", 8, "0.125", ""},
		{"negative amount", "-6.00", 2, "-3.00", ""},
		{"negative divisor", "6.00", -2, "-3.00", ""},
		{"both negative", "-6.00", -2, "3.00", ""},
		{"divide by one", "42.99", 1, "42.99", ""},
		{"zero dividend", "0.00", 5, "0.00", ""},
		{"by zero", "1.00", 0, "", "division by zero"},
		{"non-terminating limit", "1.00", 7, "", "non-terminating"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := mustParse(t, tc.in, mint.EUR)
			got, err := m.Div(tc.n)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.wantErr == "division by zero" && !errors.Is(err, mint.ErrDivisionByZero) {
					t.Errorf("err = %v, want ErrDivisionByZero", err)
				}
				if tc.wantErr == "non-terminating" && !strings.Contains(err.Error(), "non-terminating") {
					t.Errorf("err = %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Div: %v", err)
			}
			assertString(t, got, tc.want)
		})
	}
}

func TestDivImmutability(t *testing.T) {
	m := mustParse(t, "10.00", mint.EUR)
	before := m.String()
	_, _ = m.Div(2)
	if m.String() != before {
		t.Error("Div should not mutate receiver")
	}
}
