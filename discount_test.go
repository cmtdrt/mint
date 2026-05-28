package mint_test

import (
	"errors"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestDiscountAmount(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	disc := mustParse(t, "15.50", mint.EUR)
	got, err := m.DiscountAmount(disc)
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, got, "84.50")
}

func TestDiscountAmountCurrencyMismatch(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	disc := mustParse(t, "10.00", mint.USD)
	_, err := m.DiscountAmount(disc)
	if !errors.Is(err, mint.ErrCurrencyMismatch) {
		t.Errorf("err = %v, want ErrCurrencyMismatch", err)
	}
}

