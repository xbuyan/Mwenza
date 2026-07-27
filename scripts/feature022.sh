#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 022 - Complete Quantity"
echo "======================================"

cat > internal/platform/shared/quantity/quantity.go <<'EOGO'
package quantity

import "errors"

var (
	ErrNegativeQuantity = errors.New("quantity cannot be negative")
	ErrInsufficientQuantity = errors.New("insufficient quantity")
)

type Quantity struct {
	value int64
}

func New(v int64) (Quantity, error) {
	if v < 0 {
		return Quantity{}, ErrNegativeQuantity
	}

	return Quantity{value: v}, nil
}

func (q Quantity) Value() int64 {
	return q.value
}

func (q Quantity) Add(other Quantity) Quantity {
	return Quantity{
		value: q.value + other.value,
	}
}

func (q Quantity) Subtract(other Quantity) (Quantity, error) {
	if other.value > q.value {
		return Quantity{}, ErrInsufficientQuantity
	}

	return Quantity{
		value: q.value - other.value,
	}, nil
}

func (q Quantity) IsZero() bool {
	return q.value == 0
}
EOGO

cat > internal/platform/shared/quantity/quantity_test.go <<'EOGO'
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
EOGO

gofmt -w internal/platform/shared/quantity

go test ./internal/platform/shared/quantity

echo
echo "Feature 022 completed successfully."
