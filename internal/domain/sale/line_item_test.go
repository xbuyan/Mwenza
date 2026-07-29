package sale

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestCreateLineItem(t *testing.T) {
	qty, _ := quantity.New(10)

	price := money.New(750, currency.KES)

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
	price := money.New(100, currency.KES)

	_, err := NewLineItem("", "Item", qty, price)

	if err != ErrEmptyProductID {
		t.Fatal("expected empty product id error")
	}
}
