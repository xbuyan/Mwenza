#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 024 - Receive Stock"
echo "======================================"

########################################
# receive_stock.go
########################################

cat > internal/domain/inventory/receive_stock.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidReceiveQuantity = errors.New("received quantity must be greater than zero")

func (i *Inventory) ReceiveStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReceiveQuantity
	}

	i.onHand = i.onHand.Add(q)

	return nil
}
EOGO

########################################
# receive_stock_test.go
########################################

cat > internal/domain/inventory/receive_stock_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReceiveStock(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	q, _ := quantity.New(50)

	if err := inv.ReceiveStock(q); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 50 {
		t.Fatalf("expected on hand 50, got %d", inv.OnHand().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 50 {
		t.Fatalf("expected available 50, got %d", available.Value())
	}
}

func TestReceiveZeroStock(t *testing.T) {
	inv, _ := New("prod-001")

	q, _ := quantity.New(0)

	err := inv.ReceiveStock(q)

	if err != ErrInvalidReceiveQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidReceiveQuantity, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 024 completed successfully."
