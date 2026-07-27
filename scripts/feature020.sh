#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 020 - Inventory Constructor"
echo "======================================"

########################################
# inventory.go
########################################

cat > internal/domain/inventory/inventory.go <<'EOGO'
package inventory

import "errors"

var ErrEmptyProductID = errors.New("product id cannot be empty")

type Inventory struct {
	productID string
	onHand    int
	reserved  int
}

func New(productID string) (*Inventory, error) {
	if productID == "" {
		return nil, ErrEmptyProductID
	}

	return &Inventory{
		productID: productID,
		onHand:    0,
		reserved:  0,
	}, nil
}
EOGO

########################################
# inventory_test.go
########################################

cat > internal/domain/inventory/inventory_test.go <<'EOGO'
package inventory

import "testing"

func TestCreateInventory(t *testing.T) {
	inv, err := New("prod-001")

	if err != nil {
		t.Fatal(err)
	}

	if inv.onHand != 0 {
		t.Fatalf("expected zero on hand")
	}

	if inv.reserved != 0 {
		t.Fatalf("expected zero reserved")
	}
}

func TestCreateInventoryWithoutProduct(t *testing.T) {
	_, err := New("")

	if err != ErrEmptyProductID {
		t.Fatalf("expected %v got %v", ErrEmptyProductID, err)
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 020 completed successfully."
