package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestSubtract(t *testing.T) {
	a := New(500, currency.KES)
	b := New(200, currency.KES)

	c, err := a.Subtract(b)
	if err != nil {
		t.Fatal(err)
	}

	if c.Amount() != 300 {
		t.Fatalf("expected 300 got %d", c.Amount())
	}
}

func TestSubtractCurrencyMismatch(t *testing.T) {
	a := New(500, currency.KES)
	b := New(200, currency.USD)

	_, err := a.Subtract(b)

	if err != ErrCurrencyMismatch {
		t.Fatalf("expected %v got %v", ErrCurrencyMismatch, err)
	}
}
