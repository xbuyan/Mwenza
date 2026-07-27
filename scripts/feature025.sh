#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 025 - Reserve Stock"
echo "======================================"

########################################
# reserve_stock.go
########################################

cat > internal/domain/inventory/reserve_stock.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrInvalidReservationQuantity = errors.New("reservation quantity must be greater than zero")
	ErrInsufficientAvailableStock = errors.New("insufficient available stock")
)

func (i *Inventory) ReserveStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReservationQuantity
	}

	available, err := i.Available()
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	remaining, err := available.Subtract(q)
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	_ = remaining // confirms the subtraction succeeded

	i.reserved = i.reserved.Add(q)

	return nil
}
EOGO

########################################
# reserve_stock_test.go
########################################

cat > internal/domain/inventory/reserve_stock_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReserveStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserve, _ := quantity.New(30)

	if err := inv.ReserveStock(reserve); err != nil {
		t.Fatal(err)
	}

	if inv.Reserved().Value() != 30 {
		t.Fatalf("expected reserved 30, got %d", inv.Reserved().Value())
	}

	available, _ := inv.Available()

	if available.Value() != 70 {
		t.Fatalf("expected available 70, got %d", available.Value())
	}
}

func TestReserveTooMuchStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(10)
	_ = inv.ReceiveStock(received)

	reserve, _ := quantity.New(20)

	err := inv.ReserveStock(reserve)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientAvailableStock, err)
	}
}

func TestReserveZeroStock(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.ReserveStock(zero)

	if err != ErrInvalidReservationQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidReservationQuantity, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 025 completed successfully."
