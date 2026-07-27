package product

import "testing"

func TestChangeDescription(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.ChangeDescription("Premium Portland Cement")

	if p.Description() != "Premium Portland Cement" {
		t.Fatalf("expected description to be updated")
	}
}

func TestEmptyDescriptionIsAllowed(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	p.ChangeDescription("")

	if p.Description() != "" {
		t.Fatalf("expected empty description")
	}
}
