#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 014 - Change Product Unit"
echo "======================================"

########################################
# change_unit.go
########################################

cat > internal/domain/product/change_unit.go <<'EOGO'
package product

func (p *Product) ChangeUnit(unit Unit) error {
	if unit == "" {
		return ErrInvalidUnit
	}

	if p.unit == unit {
		return nil
	}

	p.unit = unit
	return nil
}
EOGO

########################################
# change_unit_test.go
########################################

cat > internal/domain/product/change_unit_test.go <<'EOGO'
package product

import "testing"

func TestChangeUnit(t *testing.T) {
	p, err := New("prod-001", "CEM-001", "Cement", UnitBag)
	if err != nil {
		t.Fatal(err)
	}

	err = p.ChangeUnit(UnitPiece)
	if err != nil {
		t.Fatal(err)
	}

	if p.Unit() != UnitPiece {
		t.Fatalf("expected %s, got %s", UnitPiece, p.Unit())
	}
}

func TestChangeUnitEmpty(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	err := p.ChangeUnit("")

	if err != ErrInvalidUnit {
		t.Fatalf("expected %v, got %v", ErrInvalidUnit, err)
	}
}

func TestChangeUnitSameValue(t *testing.T) {
	p, _ := New("prod-001", "CEM-001", "Cement", UnitBag)

	err := p.ChangeUnit(UnitBag)
	if err != nil {
		t.Fatal(err)
	}

	if p.Unit() != UnitBag {
		t.Fatalf("expected %s, got %s", UnitBag, p.Unit())
	}
}
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 014 completed successfully."
