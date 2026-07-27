#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 019 - Inventory Skeleton"
echo "======================================"

mkdir -p internal/domain/inventory

########################################
# inventory.go
########################################

cat > internal/domain/inventory/inventory.go <<'EOGO'
package inventory

type Inventory struct {
	productID string
	onHand    int
}
EOGO

########################################
# inventory_test.go
########################################

cat > internal/domain/inventory/inventory_test.go <<'EOGO'
package inventory

import "testing"

func TestInventoryCanBeCreated(t *testing.T) {
	i := Inventory{}

	if i.onHand != 0 {
		t.Fatalf("expected zero stock")
	}
}
EOGO

gofmt -w internal/domain/inventory

go test ./internal/domain/inventory

echo
echo "Feature 019 completed successfully."
