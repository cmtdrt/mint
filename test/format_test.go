package test

import (
	"math/big"
	"testing"

	"github.com/cmtdrt/mint"
)

func TestFormat(t *testing.T) {
	m := mustParse(t, "10.99", mint.EUR)
	f := m.Format()
	if f.Amount != "10.99" {
		t.Errorf("Amount = %q", f.Amount)
	}
	if f.Currency != mint.EUR {
		t.Errorf("Currency = %v", f.Currency)
	}
}

func TestFormatScaleZero(t *testing.T) {
	m := mustParse(t, "42", mint.USD)
	f := m.Format()
	if f.Amount != "42" || f.Currency != mint.USD {
		t.Errorf("got %+v", f)
	}
}

func TestFormatSmallScalePadding(t *testing.T) {
	m := mint.New(big.NewInt(3), mint.EUR, 3)
	f := m.Format()
	if f.Amount != "0.003" {
		t.Errorf("Amount = %q, want 0.003", f.Amount)
	}
}
