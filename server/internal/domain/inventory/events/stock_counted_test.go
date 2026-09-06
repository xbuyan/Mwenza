package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestNewStockCounted(t *testing.T) {
	productID := ids.New()

	qty, err := quantity.New(73)
	if err != nil {
		t.Fatal(err)
	}

	event := NewStockCounted(
		productID,
		qty,
	)

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

	if event.EventName() != StockCountedEventName {
		t.Fatalf(
			"expected event name %q, got %q",
			StockCountedEventName,
			event.EventName(),
		)
	}

	if event.EventID().IsZero() {
		t.Fatal("expected event ID")
	}

	if event.OccurredAt().IsZero() {
		t.Fatal("expected occurrence time")
	}
}
