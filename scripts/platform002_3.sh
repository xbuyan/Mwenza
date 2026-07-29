#!/usr/bin/env bash
set -e

echo "======================================"
echo "Platform 002.3 - Sale IDs"
echo "======================================"

########################################
# sale.go
########################################

cat > internal/domain/sale/sale.go <<'EOGO'
package sale

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

var ErrEmptySaleID = errors.New("sale id cannot be empty")

type Sale struct {
	id     ids.ID
	status Status
}

func New(id ids.ID) (*Sale, error) {
	if id.IsZero() {
		return nil, ErrEmptySaleID
	}

	return &Sale{
		id:     id,
		status: StatusDraft,
	}, nil
}
EOGO

########################################
# line_item.go
########################################

cat > internal/domain/sale/line_item.go <<'EOGO'
package sale

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrEmptyProductID   = errors.New("product id cannot be empty")
	ErrEmptyProductName = errors.New("product name cannot be empty")
)

type LineItem struct {
	productID   ids.ID
	productName string
	quantity    quantity.Quantity
	unitPrice   money.Money
}

func NewLineItem(
	productID ids.ID,
	productName string,
	qty quantity.Quantity,
	price money.Money,
) (LineItem, error) {

	if productID.IsZero() {
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
# sale_test.go
########################################

cat > internal/domain/sale/sale_test.go <<'EOGO'
package sale

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

func TestCreateSale(t *testing.T) {
	id := ids.New()

	s, err := New(id)
	if err != nil {
		t.Fatal(err)
	}

	if s.id != id {
		t.Fatalf("unexpected sale id")
	}

	if s.status != StatusDraft {
		t.Fatalf("expected draft status")
	}
}

func TestCreateSaleWithoutID(t *testing.T) {
	var id ids.ID

	_, err := New(id)

	if err != ErrEmptySaleID {
		t.Fatalf("expected %v got %v", ErrEmptySaleID, err)
	}
}
EOGO

########################################
# line_item_test.go
########################################

cat > internal/domain/sale/line_item_test.go <<'EOGO'
package sale

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/currency"
	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestCreateLineItem(t *testing.T) {
	qty, _ := quantity.New(10)
	price := money.New(750, currency.KES)
	productID := ids.New()

	item, err := NewLineItem(
		productID,
		"Bamburi Cement 50kg",
		qty,
		price,
	)

	if err != nil {
		t.Fatal(err)
	}

	if item.productID != productID {
		t.Fatal("unexpected product id")
	}

	if item.productName != "Bamburi Cement 50kg" {
		t.Fatal("unexpected product name")
	}
}

func TestEmptyProductID(t *testing.T) {
	qty, _ := quantity.New(1)
	price := money.New(100, currency.KES)

	var productID ids.ID

	_, err := NewLineItem(productID, "Item", qty, price)

	if err != ErrEmptyProductID {
		t.Fatal("expected empty product id error")
	}
}
EOGO

gofmt -w internal/domain/sale

echo
echo "Platform 002.3 migration applied."
