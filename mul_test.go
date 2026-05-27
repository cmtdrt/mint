package mint_test

import (
	"testing"

	"github.com/cmtdrt/mint"
)

func TestMul(t *testing.T) {
	cases := []struct {
		name string
		in   string
		n    int64
		want string
	}{
		{"double", "10.92", 2, "21.84"},
		{"triple", "10.92", 3, "32.76"},
		{"by one", "99.99", 1, "99.99"},
		{"by zero", "123.45", 0, "0.00"},
		{"integer amount", "100", 5, "500"},
		{"negative factor", "4.00", -2, "-8.00"},
		{"negative amount", "-3.50", 2, "-7.00"},
		{"both negative", "-2.00", -3, "6.00"},
		{"scale preserved small", "0.01", 10, "0.10"},
		{"large factor", "1.00", 1000, "1000.00"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := mustParse(t, tc.in, mint.EUR)
			got := m.Mul(tc.n)
			assertString(t, got, tc.want)
		})
	}
}

func TestMulImmutability(t *testing.T) {
	m := mustParse(t, "2.00", mint.EUR)
	before := m.String()
	_ = m.Mul(5)
	if m.String() != before {
		t.Error("Mul should not mutate receiver")
	}
}

func TestMulPreservesScale(t *testing.T) {
	m := mustParse(t, "1.234", mint.EUR)
	if m.Scale != 3 {
		t.Fatalf("setup scale = %d, want 3", m.Scale)
	}
	got := m.Mul(2)
	if got.Scale != 3 {
		t.Errorf("Scale = %d, want 3", got.Scale)
	}
	assertString(t, got, "2.468")
}
