package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestEqual(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.KES)

	if !a.Equal(b) {
		t.Fatal("expected monies to be equal")
	}
}

func TestEqualDifferentAmount(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	if a.Equal(b) {
		t.Fatal("expected monies to be different")
	}
}

func TestEqualDifferentCurrency(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.USD)

	if a.Equal(b) {
		t.Fatal("expected monies to be different")
	}
}
