#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 027 - Dispatch Reserved Stock"
echo "======================================"

########################################
# dispatch_reserved_stock.go
########################################

cat > internal/domain/inventory/dispatch_reserved_stock.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDispatchQuantity = errors.New("dispatch quantity must be greater than zero")

func (i *Inventory) DispatchReservedStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidDispatchQuantity
	}

	reservedRemaining, err := i.reserved.Subtract(q)
	if err != nil {
		return ErrInsufficientReservedStock
	}

	onHandRemaining, err := i.onHand.Subtract(q)
	if err != nil {
		// This should never happen if inventory invariants are respected.
		return err
	}

	i.reserved = reservedRemaining
	i.onHand = onHandRemaining

	return nil
}
EOGO

########################################
# dispatch_reserved_stock_test.go
########################################

cat > internal/domain/inventory/dispatch_reserved_stock_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestDispatchReservedStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(40)
	_ = inv.ReserveStock(reserved)

	dispatch, _ := quantity.New(25)

	if err := inv.DispatchReservedStock(dispatch); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 75 {
		t.Fatalf("expected on hand 75, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 15 {
		t.Fatalf("expected reserved 15, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 60 {
		t.Fatalf("expected available 60, got %d", available.Value())
	}
}

func TestDispatchMoreThanReserved(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(50)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(10)
	_ = inv.ReserveStock(reserved)

	dispatch, _ := quantity.New(20)

	err := inv.DispatchReservedStock(dispatch)

	if err != ErrInsufficientReservedStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientReservedStock, err)
	}
}

func TestDispatchZeroQuantity(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.DispatchReservedStock(zero)

	if err != ErrInvalidDispatchQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidDispatchQuantity, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 027 completed successfully."
