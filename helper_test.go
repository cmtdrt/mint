package mint_test

import (
	"testing"

	"github.com/cmtdrt/mint"
)

func mustParse(t *testing.T, input string, currency mint.Currency) mint.Mint {
	t.Helper()
	m, err := mint.NewFromStr(input, currency)
	if err != nil {
		t.Fatalf("NewFromStr(%q): %v", input, err)
	}
	return m
}

func assertString(t *testing.T, m mint.Mint, want string) {
	t.Helper()
	if got := m.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
