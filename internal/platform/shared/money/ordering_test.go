package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestOrdering(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	gt, _ := b.GreaterThan(a)
	if !gt {
		t.Fatal("expected greater")
	}

	lt, _ := a.LessThan(b)
	if !lt {
		t.Fatal("expected less")
	}

	ge, _ := b.GreaterOrEqual(b)
	if !ge {
		t.Fatal("expected greater or equal")
	}

	le, _ := a.LessOrEqual(a)
	if !le {
		t.Fatal("expected less or equal")
	}
}
