package quantity

import "testing"

func TestNegativeQuantity(t *testing.T) {
	_, err := New(-1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAddQuantity(t *testing.T) {
	a, _ := New(10)
	b, _ := New(5)

	c := a.Add(b)

	if c.Value() != 15 {
		t.Fatalf("expected 15, got %d", c.Value())
	}
}

func TestSubtractQuantity(t *testing.T) {
	a, _ := New(10)
	b, _ := New(4)

	c, err := a.Subtract(b)
	if err != nil {
		t.Fatal(err)
	}

	if c.Value() != 6 {
		t.Fatalf("expected 6, got %d", c.Value())
	}
}

func TestSubtractTooMuch(t *testing.T) {
	a, _ := New(5)
	b, _ := New(10)

	_, err := a.Subtract(b)

	if err != ErrInsufficientQuantity {
		t.Fatalf("expected %v, got %v", ErrInsufficientQuantity, err)
	}
}

func TestZeroQuantity(t *testing.T) {
	q, _ := New(0)

	if !q.IsZero() {
		t.Fatal("expected zero quantity")
	}
}
