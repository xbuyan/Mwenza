#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 016 - Prevent Reactivation"
echo "======================================"

########################################
# lifecycle.go
########################################

cat > internal/domain/product/lifecycle.go <<'EOGO'
package product

func (p *Product) Activate() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	if p.status == StatusInactive {
		p.status = StatusActive
	}
}

func (p *Product) Deactivate() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	if p.status == StatusActive {
		p.status = StatusInactive
	}
}
EOGO

########################################
# lifecycle_test.go
########################################

cat > internal/domain/product/lifecycle_test.go <<'EOGO'
package product

import "testing"

func TestDeactivateProduct(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.Deactivate()

	if p.Status() != StatusInactive {
		t.Fatalf("expected inactive, got %s", p.Status())
	}
}

func TestActivateProduct(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.Deactivate()
	p.Activate()

	if p.Status() != StatusActive {
		t.Fatalf("expected active, got %s", p.Status())
	}
}

func TestCannotActivateDiscontinuedProduct(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.Discontinue()
	p.Activate()

	if p.Status() != StatusDiscontinued {
		t.Fatalf("expected %s, got %s", StatusDiscontinued, p.Status())
	}
}

func TestCannotDeactivateDiscontinuedProduct(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.Discontinue()
	p.Deactivate()

	if p.Status() != StatusDiscontinued {
		t.Fatalf("expected %s, got %s", StatusDiscontinued, p.Status())
	}
}
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 016 completed successfully."
