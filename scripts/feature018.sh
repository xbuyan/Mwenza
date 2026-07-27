#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 018 - Product Description"
echo "======================================"

########################################
# product.go
########################################

cat > internal/domain/product/product.go <<'EOGO'
package product

type Product struct {
	id          string
	sku         SKU
	name        string
	description string
	unit        Unit
	status      Status
}

func New(id string, sku SKU, name string, unit Unit) (*Product, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	if sku == "" {
		return nil, ErrEmptySKU
	}

	if name == "" {
		return nil, ErrEmptyName
	}

	if unit == "" {
		return nil, ErrInvalidUnit
	}

	return &Product{
		id:          id,
		sku:         sku,
		name:        name,
		description: "",
		unit:        unit,
		status:      StatusActive,
	}, nil
}
EOGO

########################################
# description.go
########################################

cat > internal/domain/product/description.go <<'EOGO'
package product

func (p *Product) Description() string {
	return p.description
}

func (p *Product) ChangeDescription(description string) {
	if p == nil {
		return
	}

	p.description = description
}
EOGO

########################################
# description_test.go
########################################

cat > internal/domain/product/description_test.go <<'EOGO'
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
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 018 completed successfully."
