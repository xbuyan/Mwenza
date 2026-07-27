package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestAdd(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	c, err := a.Add(b)
	if err != nil {
		t.Fatal(err)
	}

	if c.Amount() != 300 {
		t.Fatalf("expected 300 got %d", c.Amount())
	}
}
