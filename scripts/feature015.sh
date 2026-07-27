#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 015 - Discontinue Product"
echo "======================================"

########################################
# discontinue.go
########################################

cat > internal/domain/product/discontinue.go <<'EOGO'
package product

func (p *Product) Discontinue() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	p.status = StatusDiscontinued
}
EOGO

########################################
# discontinue_test.go
########################################

cat > internal/domain/product/discontinue_test.go <<'EOGO'
package product

import "testing"

func TestDiscontinueProduct(t *testing.T) {
	p, err := New("prod-001", "CEM-001", "Cement", UnitBag)
	if err != nil {
		t.Fatal(err)
	}

	p.Discontinue()

	if p.Status() != StatusDiscontinued {
		t.Fatalf("expected %s, got %s", StatusDiscontinued, p.Status())
	}
}

func TestDiscontinueIsIdempotent(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.Discontinue()
	p.Discontinue()

	if p.Status() != StatusDiscontinued {
		t.Fatalf("expected %s, got %s", StatusDiscontinued, p.Status())
	}
}
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 015 completed successfully."
