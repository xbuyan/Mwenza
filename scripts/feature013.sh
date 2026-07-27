#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 013 - Rename Product"
echo "======================================"

########################################
# rename.go
########################################

cat > internal/domain/product/rename.go <<'EOGO'
package product

func (p *Product) Rename(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	p.name = name
	return nil
}
EOGO

########################################
# rename_test.go
########################################

cat > internal/domain/product/rename_test.go <<'EOGO'
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
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 013 completed successfully."
