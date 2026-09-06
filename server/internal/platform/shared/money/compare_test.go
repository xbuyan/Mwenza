package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestCompareEqual(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.KES)

	got, err := a.Compare(b)
	if err != nil {
		t.Fatal(err)
	}

	if got != 0 {
		t.Fatalf("expected 0 got %d", got)
	}
}

func TestCompareLess(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	got, err := a.Compare(b)
	if err != nil {
		t.Fatal(err)
	}

	if got != -1 {
		t.Fatalf("expected -1 got %d", got)
	}
}

func TestCompareGreater(t *testing.T) {
	a := New(300, currency.KES)
	b := New(200, currency.KES)

	got, err := a.Compare(b)
	if err != nil {
		t.Fatal(err)
	}

	if got != 1 {
		t.Fatalf("expected 1 got %d", got)
	}
}

func TestCompareCurrencyMismatch(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.USD)

	_, err := a.Compare(b)

	if err != ErrCurrencyMismatch {
		t.Fatalf("expected %v got %v", ErrCurrencyMismatch, err)
	}
}
