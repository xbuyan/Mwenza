#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 021 - Inventory Getters"
echo "======================================"

########################################
# getters.go
########################################

cat > internal/domain/inventory/getters.go <<'EOGO'
package inventory

func (i *Inventory) ProductID() string {
	return i.productID
}

func (i *Inventory) OnHand() int {
	return i.onHand
}

func (i *Inventory) Reserved() int {
	return i.reserved
}

func (i *Inventory) Available() int {
	return i.onHand - i.reserved
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

	if inv.OnHand() != 0 {
		t.Fatalf("expected onHand = 0")
	}

	if inv.Reserved() != 0 {
		t.Fatalf("expected reserved = 0")
	}

	if inv.Available() != 0 {
		t.Fatalf("expected available = 0")
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 021 completed successfully."
