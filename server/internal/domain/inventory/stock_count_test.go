package inventory

import (
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockCount(t *testing.T) {
	productID := ids.New()

	inv, err := New(productID)
	if err != nil {
		t.Fatal(err)
	}

	start, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(start); err != nil {
		t.Fatal(err)
	}

	// Remove the receive event so we can inspect only the stock count event.
	inv.Pull()

	counted, err := quantity.New(92)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.StockCount(counted); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 92 {
		t.Fatalf("expected on hand 92, got %d", inv.OnHand().Value())
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockCounted)
	if !ok {
		t.Fatalf("expected StockCounted event, got %T", events[0])
	}

	if event.ProductID != productID {
		t.Fatalf(
			"expected product ID %v, got %v",
			productID,
			event.ProductID,
		)
	}

	if !event.Quantity.Equal(counted) {
		t.Fatalf(
			"expected quantity %d, got %d",
			counted.Value(),
			event.Quantity.Value(),
		)
	}
}

func TestStockCountCannotGoBelowReserved(t *testing.T) {
	productID := ids.New()

	inv, err := New(productID)
	if err != nil {
		t.Fatal(err)
	}

	start, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(start); err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	// Remove events from setup operations.
	inv.Pull()

	counted, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.StockCount(counted)

	if err != ErrCountBelowReserved {
		t.Fatalf("expected %v got %v", ErrCountBelowReserved, err)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf(
			"expected on hand to remain 100, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.Reserved().Value() != 30 {
		t.Fatalf(
			"expected reserved to remain 30, got %d",
			inv.Reserved().Value(),
		)
	}

	if inv.HasEvents() {
		t.Fatal("expected no event for rejected stock count")
	}
}
