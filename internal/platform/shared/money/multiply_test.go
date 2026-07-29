package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestMultiply(t *testing.T) {
	m := New(250, currency.KES)

	got := m.Multiply(4)

	if got.Amount() != 1000 {
		t.Fatalf("expected 1000 got %d", got.Amount())
	}

	if got.Currency() != currency.KES {
		t.Fatal("currency changed")
	}
}
