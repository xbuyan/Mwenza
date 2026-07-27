#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 023 - Inventory Uses Quantity"
echo "======================================"

########################################
# inventory.go
########################################

cat > internal/domain/inventory/inventory.go <<'EOGO'
package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrEmptyProductID = errors.New("product id cannot be empty")

type Inventory struct {
	productID string
	onHand    quantity.Quantity
	reserved  quantity.Quantity
}

func New(productID string) (*Inventory, error) {
	if productID == "" {
		return nil, ErrEmptyProductID
	}

	zero, err := quantity.New(0)
	if err != nil {
		return nil, err
	}

	return &Inventory{
		productID: productID,
		onHand:    zero,
		reserved:  zero,
	}, nil
}
EOGO

########################################
# getters.go
########################################

cat > internal/domain/inventory/getters.go <<'EOGO'
package inventory

import "github.com/mwenza/mwenza/internal/platform/shared/quantity"

func (i *Inventory) ProductID() string {
	return i.productID
}

func (i *Inventory) OnHand() quantity.Quantity {
	return i.onHand
}

func (i *Inventory) Reserved() quantity.Quantity {
	return i.reserved
}

func (i *Inventory) Available() (quantity.Quantity, error) {
	return i.onHand.Subtract(i.reserved)
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

	if inv.OnHand().Value() != 0 {
		t.Fatalf("expected zero on hand")
	}

	if inv.Reserved().Value() != 0 {
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

########################################
# getters_test.go
########################################

cat > internal/domain/inventory/getters_test.go <<'EOGO'
package inventory

import "testing"

func TestInventoryGetters(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	if inv.ProductID() != "prod-001" {
		t.Fatalf("unexpected product id")
	}

	if inv.OnHand().Value() != 0 {
		t.Fatalf("expected onHand = 0")
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected reserved = 0")
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 0 {
		t.Fatalf("expected available = 0")
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 023 completed successfully."
