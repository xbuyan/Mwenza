package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockReserved(t *testing.T) {
	productID := ids.New()

	qty, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	event := NewStockReserved(productID, qty)

	if event.EventName() != StockReservedEventName {
		t.Fatalf(
			"expected event name %q, got %q",
			StockReservedEventName,
			event.EventName(),
		)
	}

	if event.ProductID != productID {
		t.Fatalf(
			"expected product ID %v, got %v",
			productID,
			event.ProductID,
		)
	}

	if !event.Quantity.Equal(qty) {
		t.Fatalf(
			"expected quantity %d, got %d",
			qty.Value(),
			event.Quantity.Value(),
		)
	}
}
