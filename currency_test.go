package mint_test

import (
	"testing"

	"github.com/cmtdrt/mint"
)

func TestCurrencyString(t *testing.T) {
	if mint.EUR.String() != "EUR" {
		t.Errorf("EUR.String() = %q", mint.EUR.String())
	}
	if mint.USD.String() != "USD" {
		t.Errorf("USD.String() = %q", mint.USD.String())
	}
}

func TestFromString(t *testing.T) {
	c, err := mint.FromString("eur")
	if err != nil || c != mint.EUR {
		t.Fatalf("FromString(eur) = %v, %v", c, err)
	}
	_, err = mint.FromString("GBP")
	if err == nil {
		t.Fatal("expected error for GBP")
	}
}
