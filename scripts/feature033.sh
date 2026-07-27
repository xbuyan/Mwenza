#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 033 - Line Item Value Object"
echo "======================================"

mkdir -p internal/domain/sale

########################################
# line_item.go
########################################

cat > internal/domain/sale/line_item.go <<'EOGO'
package sale

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrEmptyProductID   = errors.New("product id cannot be empty")
	ErrEmptyProductName = errors.New("product name cannot be empty")
)

type LineItem struct {
	productID   string
	productName string
	quantity    quantity.Quantity
	unitPrice   money.Money
}

func NewLineItem(
	productID string,
	productName string,
	qty quantity.Quantity,
	price money.Money,
) (LineItem, error) {

	if productID == "" {
		return LineItem{}, ErrEmptyProductID
	}

	if productName == "" {
		return LineItem{}, ErrEmptyProductName
	}

	return LineItem{
		productID:   productID,
		productName: productName,
		quantity:    qty,
		unitPrice:   price,
	}, nil
}
EOGO

########################################
# line_item_test.go
########################################

cat > internal/domain/sale/line_item_test.go <<'EOGO'
package sale

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestCreateLineItem(t *testing.T) {
	qty, _ := quantity.New(10)

	price, _ := money.New(750, currency.KES)

	item, err := NewLineItem(
		"prod-001",
		"Bamburi Cement 50kg",
		qty,
		price,
	)

	if err != nil {
		t.Fatal(err)
	}

	if item.productID != "prod-001" {
		t.Fatal("unexpected product id")
	}

	if item.productName != "Bamburi Cement 50kg" {
		t.Fatal("unexpected product name")
	}
}

func TestEmptyProductID(t *testing.T) {
	qty, _ := quantity.New(1)
	price, _ := money.New(100, currency.KES)

	_, err := NewLineItem("", "Item", qty, price)

	if err != ErrEmptyProductID {
		t.Fatal("expected empty product id error")
	}
}
EOGO

gofmt -w internal/domain/sale

go test ./internal/domain/sale

echo
echo "Feature 033 completed successfully."
