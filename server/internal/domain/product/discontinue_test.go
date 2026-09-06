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
