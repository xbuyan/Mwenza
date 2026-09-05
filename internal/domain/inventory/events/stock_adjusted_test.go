package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestNewStockAdjusted(t *testing.T) {
	productID := ids.New()

	qty, err := quantity.New(25)
	if err != nil {
		t.Fatal(err)
	}

	event := NewStockAdjusted(
		productID,
		qty,
		1,
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

	if event.Direction != 1 {
		t.Fatalf(
			"expected direction 1, got %d",
			event.Direction,
		)
	}

	if event.EventName() != StockAdjustedEventName {
		t.Fatalf(
			"expected event name %q, got %q",
			StockAdjustedEventName,
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
