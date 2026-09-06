package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockDispatched(t *testing.T) {
	productID := ids.New()

	qty, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	event := NewStockDispatched(productID, qty)

	if event.EventID().IsZero() {
		t.Fatal("expected event ID to be non-zero")
	}

	if event.OccurredAt().IsZero() {
		t.Fatal("expected occurred-at timestamp to be non-zero")
	}

	if event.EventName() != StockDispatchedEventName {
		t.Fatalf(
			"expected event name %q, got %q",
			StockDispatchedEventName,
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
