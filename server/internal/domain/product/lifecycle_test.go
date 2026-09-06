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
