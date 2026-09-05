package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockReservationReleased(t *testing.T) {
	productID := ids.New()

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatal(err)
	}

	event := NewStockReservationReleased(productID, qty)

	if event.EventName() != StockReservationReleasedEventName {
		t.Fatalf(
			"expected event name %q, got %q",
			StockReservationReleasedEventName,
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
