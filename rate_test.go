package mint_test

import (
	"errors"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestTaxRateParsingHappyPath(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)

	tax, err := m.Tax("0.20")
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, tax, "20.0000")
}

func TestTaxInvalidRate(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	_, err := m.Tax("abc")
	if !errors.Is(err, mint.ErrInvalidRate) {
		t.Errorf("err = %v, want ErrInvalidRate", err)
	}
}

func TestTaxNegativeRateRejected(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	_, err := m.Tax("-0.20")
	if !errors.Is(err, mint.ErrInvalidOperation) {
		t.Errorf("err = %v, want ErrInvalidOperation", err)
	}
}

func TestDiscountPercentUsesRatio(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	got, err := m.DiscountPercent("0.20")
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, got, "80.0000")
}

func TestDiscountPercentOverOneRejected(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	_, err := m.DiscountPercent("1.01")
	if !errors.Is(err, mint.ErrInvalidOperation) {
		t.Errorf("err = %v, want ErrInvalidOperation", err)
	}
}
