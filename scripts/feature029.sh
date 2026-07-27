#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 029 - Inventory Adjustment"
echo "======================================"

########################################
# adjust_stock.go
########################################

cat > internal/domain/inventory/adjust_stock.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type AdjustmentDirection int

const (
	Increase AdjustmentDirection = iota
	Decrease
)

var (
	ErrInvalidAdjustmentQuantity = errors.New("adjustment quantity must be greater than zero")
	ErrAdjustmentBelowReserved   = errors.New("adjustment would reduce stock below reserved quantity")
)

func (i *Inventory) AdjustStock(direction AdjustmentDirection, q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidAdjustmentQuantity
	}

	switch direction {
	case Increase:
		i.onHand = i.onHand.Add(q)
		return nil

	case Decrease:
		newOnHand, err := i.onHand.Subtract(q)
		if err != nil {
			return ErrAdjustmentBelowReserved
		}

		if _, err := newOnHand.Subtract(i.reserved); err != nil {
			return ErrAdjustmentBelowReserved
		}

		i.onHand = newOnHand
		return nil

	default:
		return errors.New("invalid adjustment direction")
	}
}
EOGO

########################################
# adjust_stock_test.go
########################################

cat > internal/domain/inventory/adjust_stock_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestIncreaseStock(t *testing.T) {
	inv, _ := New("prod-001")

	q, _ := quantity.New(20)

	if err := inv.AdjustStock(Increase, q); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf("expected on hand 20")
	}
}

func TestDecreaseStock(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(50)
	_ = inv.ReceiveStock(start)

	dec, _ := quantity.New(10)

	if err := inv.AdjustStock(Decrease, dec); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 40 {
		t.Fatalf("expected on hand 40")
	}
}

func TestCannotAdjustBelowReserved(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(20)
	_ = inv.ReceiveStock(start)

	reserved, _ := quantity.New(15)
	_ = inv.ReserveStock(reserved)

	dec, _ := quantity.New(10)

	err := inv.AdjustStock(Decrease, dec)

	if err != ErrAdjustmentBelowReserved {
		t.Fatalf("expected %v got %v", ErrAdjustmentBelowReserved, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 029 completed successfully."
