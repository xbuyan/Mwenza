package quantity

import "testing"

func TestCompare(t *testing.T) {
	a, _ := New(10)
	b, _ := New(20)
	c, _ := New(10)

	if a.Compare(b) != -1 {
		t.Fatal("expected -1")
	}

	if b.Compare(a) != 1 {
		t.Fatal("expected 1")
	}

	if a.Compare(c) != 0 {
		t.Fatal("expected 0")
	}
}

func TestEqual(t *testing.T) {
	a, _ := New(5)
	b, _ := New(5)

	if !a.Equal(b) {
		t.Fatal("expected equal")
	}
}

func TestOrderingHelpers(t *testing.T) {
	a, _ := New(5)
	b, _ := New(10)

	if !b.GreaterThan(a) {
		t.Fatal("expected greater")
	}

	if !a.LessThan(b) {
		t.Fatal("expected less")
	}

	if !a.LessOrEqual(b) {
		t.Fatal("expected less or equal")
	}

	if !b.GreaterOrEqual(a) {
		t.Fatal("expected greater or equal")
	}
}
