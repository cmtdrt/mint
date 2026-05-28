package mint_test

import (
	"errors"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestTaxExclusive(t *testing.T) {
	m := mustParse(t, "100.00", mint.EUR)
	got, err := m.TaxExclusive("0.20")
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, got, "120.0000")
}

func TestTaxInclusiveExact(t *testing.T) {
	// gross=120 with 20% VAT => net=100
	gross := mustParse(t, "120.00", mint.EUR)
	net, err := gross.TaxInclusive("0.20")
	if err != nil {
		t.Fatal(err)
	}
	assertString(t, net, "100.00")
}

func TestTaxInclusiveNonTerminating(t *testing.T) {
	// gross / 1.07 is repeating (for most values), should error
	gross := mustParse(t, "100.00", mint.EUR)
	_, err := gross.TaxInclusive("0.07")
	if !errors.Is(err, mint.ErrNonTerminatingResult) {
		t.Errorf("err = %v, want ErrNonTerminatingResult", err)
	}
}

