#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 030 - Stock Count"
echo "======================================"

########################################
# stock_count.go
########################################

cat > internal/domain/inventory/stock_count.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrCountBelowReserved = errors.New("counted stock cannot be below reserved stock")

func (i *Inventory) StockCount(actual quantity.Quantity) error {
	if _, err := actual.Subtract(i.reserved); err != nil {
		return ErrCountBelowReserved
	}

	i.onHand = actual
	return nil
}
EOGO

########################################
# stock_count_test.go
########################################

cat > internal/domain/inventory/stock_count_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockCount(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(100)
	_ = inv.ReceiveStock(start)

	counted, _ := quantity.New(92)

	if err := inv.StockCount(counted); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 92 {
		t.Fatalf("expected on hand 92")
	}
}

func TestStockCountCannotGoBelowReserved(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(100)
	_ = inv.ReceiveStock(start)

	reserved, _ := quantity.New(30)
	_ = inv.ReserveStock(reserved)

	counted, _ := quantity.New(20)

	err := inv.StockCount(counted)

	if err != ErrCountBelowReserved {
		t.Fatalf("expected %v got %v", ErrCountBelowReserved, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 030 completed successfully."
