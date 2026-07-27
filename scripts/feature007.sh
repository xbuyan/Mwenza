#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 007 - Product Domain Skeleton"
echo "======================================"

mkdir -p internal/domain/product

########################################
# product.go
########################################

cat > internal/domain/product/product.go <<'EOGO'
package product

type Product struct {
	id     string
	sku    SKU
	name   string
	unit   Unit
	status Status
}
EOGO

########################################
# sku.go
########################################

cat > internal/domain/product/sku.go <<'EOGO'
package product

type SKU string
EOGO

########################################
# unit.go
########################################

cat > internal/domain/product/unit.go <<'EOGO'
package product

type Unit string

const (
	UnitPiece    Unit = "piece"
	UnitBag      Unit = "bag"
	UnitBox      Unit = "box"
	UnitKilogram Unit = "kilogram"
	UnitLitre    Unit = "litre"
	UnitMetre    Unit = "metre"
)
EOGO

########################################
# status.go
########################################

cat > internal/domain/product/status.go <<'EOGO'
package product

type Status string

const (
	StatusActive       Status = "active"
	StatusInactive     Status = "inactive"
	StatusDiscontinued Status = "discontinued"
)
EOGO

########################################
# errors.go
########################################

cat > internal/domain/product/errors.go <<'EOGO'
package product

import "errors"

var (
	ErrEmptyName  = errors.New("product name cannot be empty")
	ErrEmptySKU   = errors.New("product SKU cannot be empty")
	ErrInvalidUnit = errors.New("invalid product unit")
)
EOGO

gofmt -w internal/domain/product

echo
echo "Feature 007 completed successfully."
