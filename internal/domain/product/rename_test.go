package product

import "testing"

func TestRenameProduct(t *testing.T) {
	p, err := New("prod-001", "CEM-001", "Old Name", UnitBag)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Rename("New Name"); err != nil {
		t.Fatal(err)
	}

	if p.Name() != "New Name" {
		t.Fatalf("expected 'New Name', got '%s'", p.Name())
	}
}

func TestRenameProductWithEmptyName(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Old Name", UnitBag)

	err := p.Rename("")

	if err != ErrEmptyName {
		t.Fatalf("expected %v, got %v", ErrEmptyName, err)
	}
}
