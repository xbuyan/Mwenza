package events

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestNewStockReceived(t *testing.T) {
	id := ids.New()

	qty, _ := quantity.New(25)

	event := NewStockReceived(id, qty)

	if event.ProductID != id {
		t.Fatal("unexpected product id")
	}

	if !event.Quantity.Equal(qty) {
		t.Fatal("unexpected quantity")
	}

	if event.EventName() != StockReceivedEventName {
		t.Fatal("unexpected event name")
	}

	if event.EventID().IsZero() {
		t.Fatal("expected event id")
	}

	if event.OccurredAt().IsZero() {
		t.Fatal("expected occurrence time")
	}
}
