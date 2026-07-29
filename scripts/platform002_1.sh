#!/usr/bin/env bash
set -e

echo "======================================"
echo "Platform 002.1 - Product IDs"
echo "======================================"

########################################
# product.go
########################################

cat > internal/domain/product/product.go <<'EOGO'
package product

import "github.com/mwenza/mwenza/internal/platform/ids"

type Product struct {
	id          ids.ID
	sku         SKU
	name        string
	description string
	unit        Unit
	status      Status
}

func New(id ids.ID, sku SKU, name string, unit Unit) (*Product, error) {
	if id.IsZero() {
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
# getters.go
########################################

cat > internal/domain/product/getters.go <<'EOGO'
package product

import "github.com/mwenza/mwenza/internal/platform/ids"

func (p *Product) ID() ids.ID {
	return p.id
}

func (p *Product) SKU() SKU {
	return p.sku
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Unit() Unit {
	return p.unit
}

func (p *Product) Status() Status {
	return p.status
}
EOGO

gofmt -w internal/domain/product

echo
echo "Platform 002.1 migration applied."
