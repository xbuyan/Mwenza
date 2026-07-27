#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 011 - Product Equality"
echo "======================================"

########################################
# equality.go
########################################

cat > internal/domain/product/equality.go <<'EOGO'
package product

func (p *Product) Equals(other *Product) bool {
	if p == nil || other == nil {
		return false
	}

	return p.id == other.id
}
EOGO

########################################
# equality_test.go
########################################

cat > internal/domain/product/equality_test.go <<'EOGO'
package product

import "testing"

func TestProductEquals(t *testing.T) {
	p1, _ := New("prod-001", "CEM-001", "Cement", UnitBag)
	p2, _ := New("prod-001", "CEM-002", "Different Cement", UnitBag)
	p3, _ := New("prod-002", "CEM-001", "Cement", UnitBag)

	if !p1.Equals(p2) {
		t.Fatal("products with the same ID should be equal")
	}

	if p1.Equals(p3) {
		t.Fatal("products with different IDs should not be equal")
	}

	if p1.Equals(nil) {
		t.Fatal("product should not equal nil")
	}
}
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 011 completed successfully."
