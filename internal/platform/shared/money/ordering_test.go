package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestGreaterThan(t *testing.T) {
	a := New(200, currency.KES)
	b := New(100, currency.KES)

	ok, err := a.GreaterThan(b)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected greater")
	}
}

func TestLessThan(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	ok, err := a.LessThan(b)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected less")
	}
}

func TestGreaterOrEqual(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.KES)

	ok, err := a.GreaterOrEqual(b)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected greater or equal")
	}
}

func TestLessOrEqual(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.KES)

	ok, err := a.LessOrEqual(b)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected less or equal")
	}
}

func TestOrderingCurrencyMismatch(t *testing.T) {
	a := New(100, currency.KES)
	b := New(100, currency.USD)

	_, err := a.GreaterThan(b)
	if err != ErrCurrencyMismatch {
		t.Fatalf("expected %v got %v", ErrCurrencyMismatch, err)
	}
}
