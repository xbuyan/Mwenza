package product

import "testing"

func TestChangeSKU(t *testing.T) {
	p, err := New("prod-001", "CEM-001", "Cement", UnitBag)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.ChangeSKU("CEM-002"); err != nil {
		t.Fatal(err)
	}

	if p.SKU() != "CEM-002" {
		t.Fatalf("expected CEM-002, got %s", p.SKU())
	}
}

func TestChangeSKUEmpty(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	err := p.ChangeSKU("")

	if err != ErrEmptySKU {
		t.Fatalf("expected %v, got %v", ErrEmptySKU, err)
	}
}

func TestChangeSKUSameValue(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	if err := p.ChangeSKU("CEM-001"); err != nil {
		t.Fatal(err)
	}

	if p.SKU() != "CEM-001" {
		t.Fatalf("expected CEM-001, got %s", p.SKU())
	}
}
