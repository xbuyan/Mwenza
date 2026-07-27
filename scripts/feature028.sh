#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 028 - Direct Dispatch"
echo "======================================"

########################################
# dispatch_direct.go
########################################

cat > internal/domain/inventory/dispatch_direct.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDirectDispatchQuantity = errors.New("direct dispatch quantity must be greater than zero")

func (i *Inventory) DispatchDirect(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidDirectDispatchQuantity
	}

	available, err := i.Available()
	if err != nil {
		return err
	}

	remainingAvailable, err := available.Subtract(q)
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	_ = remainingAvailable

	onHandRemaining, err := i.onHand.Subtract(q)
	if err != nil {
		return err
	}

	i.onHand = onHandRemaining

	return nil
}
EOGO

########################################
# dispatch_direct_test.go
########################################

cat > internal/domain/inventory/dispatch_direct_test.go <<'EOGO'
package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestDispatchDirect(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	dispatch, _ := quantity.New(30)

	if err := inv.DispatchDirect(dispatch); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 70 {
		t.Fatalf("expected on hand 70, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected reserved 0, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 70 {
		t.Fatalf("expected available 70, got %d", available.Value())
	}
}

func TestDispatchDirectInsufficientStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(20)
	_ = inv.ReceiveStock(received)

	dispatch, _ := quantity.New(30)

	err := inv.DispatchDirect(dispatch)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientAvailableStock, err)
	}
}

func TestDispatchDirectZeroQuantity(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.DispatchDirect(zero)

	if err != ErrInvalidDirectDispatchQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidDirectDispatchQuantity, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 028 completed successfully."
