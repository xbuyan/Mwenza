package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReceiveStock(t *testing.T) {
	productID := ids.New()

	inv, err := New(productID)
	if err != nil {
		t.Fatal(err)
	}

	q, err := quantity.New(50)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(q); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 50 {
		t.Fatalf("expected on hand 50, got %d", inv.OnHand().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 50 {
		t.Fatalf("expected available 50, got %d", available.Value())
	}
}

func TestReceiveZeroStock(t *testing.T) {
	productID := ids.New()

	inv, err := New(productID)
	if err != nil {
		t.Fatal(err)
	}

	q, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.ReceiveStock(q)

	if err != ErrInvalidReceiveQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidReceiveQuantity, err)
	}
}

func TestReceiveStockRecordsEvent(t *testing.T) {
	productID := ids.New()

	inv, err := New(productID)
	if err != nil {
		t.Fatal(err)
	}

	qty, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(qty); err != nil {
		t.Fatal(err)
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if events[0].EventName() != "inventory.stock_received" {
		t.Fatalf("unexpected event %s", events[0].EventName())
	}
}
