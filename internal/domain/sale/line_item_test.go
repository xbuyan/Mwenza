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
