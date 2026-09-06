package inventory

import (
	"errors"
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestDispatchReservedStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(40)
	_ = inv.ReserveStock(reserved)

	// Clear events from receiving and reserving.
	_ = inv.Pull()

	dispatch, _ := quantity.New(25)

	if err := inv.DispatchReservedStock(dispatch); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 75 {
		t.Fatalf("expected on hand 75, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 15 {
		t.Fatalf("expected reserved 15, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 60 {
		t.Fatalf("expected available 60, got %d", available.Value())
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockDispatched)
	if !ok {
		t.Fatalf("expected StockDispatched event, got %T", events[0])
	}

	if event.ProductID != inv.ProductID() {
		t.Fatalf(
			"expected product ID %v, got %v",
			inv.ProductID(),
			event.ProductID,
		)
	}

	if event.Quantity.Value() != 25 {
		t.Fatalf(
			"expected dispatched quantity 25, got %d",
			event.Quantity.Value(),
		)
	}
}

func TestDispatchMoreThanReserved(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(50)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(10)
	_ = inv.ReserveStock(reserved)

	// Clear events from successful operations.
	_ = inv.Pull()

	dispatch, _ := quantity.New(20)

	err := inv.DispatchReservedStock(dispatch)

	if err != ErrInsufficientReservedStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientReservedStock, err)
	}

	if inv.OnHand().Value() != 50 {
		t.Fatalf("expected on hand 50, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 10 {
		t.Fatalf("expected reserved 10, got %d", inv.Reserved().Value())
	}

	if events := inv.Pull(); len(events) != 0 {
		t.Fatalf(
			"expected no events on failed dispatch, got %d",
			len(events),
		)
	}
}

func TestDispatchZeroQuantity(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.DispatchReservedStock(zero)

	if err != ErrInvalidDispatchQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidDispatchQuantity, err)
	}

	if events := inv.Pull(); len(events) != 0 {
		t.Fatalf(
			"expected no events on failed dispatch, got %d",
			len(events),
		)
	}
}

func TestDispatchReservedStockExactlyReservedQuantity(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}
	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(40)
	if err != nil {
		t.Fatal(err)
	}
	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	inv.Pull()

	if err := inv.DispatchReservedStock(reserved); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 60 {
		t.Fatalf("expected on hand 60, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected reserved 0, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 60 {
		t.Fatalf("expected available 60, got %d", available.Value())
	}

	events := inv.Pull()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockDispatched)
	if !ok {
		t.Fatalf("expected StockDispatched event, got %T", events[0])
	}

	if event.Quantity.Value() != 40 {
		t.Fatalf("expected dispatched quantity 40, got %d", event.Quantity.Value())
	}
}

func TestDispatchReservedStockRejectsZero(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(40)
	_ = inv.ReserveStock(reserved)

	inv.Pull()

	zero, _ := quantity.New(0)

	err := inv.DispatchReservedStock(zero)
	if !errors.Is(err, ErrInvalidDispatchQuantity) {
		t.Fatalf("expected ErrInvalidDispatchQuantity, got %v", err)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf("expected on hand 100, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40, got %d", inv.Reserved().Value())
	}

	if events := inv.Pull(); len(events) != 0 {
		t.Fatalf("expected no events on failed dispatch, got %d", len(events))
	}
}

func TestDispatchReservedStockRejectsInsufficientReservedStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(40)
	_ = inv.ReserveStock(reserved)

	inv.Pull()

	dispatch, _ := quantity.New(41)

	err := inv.DispatchReservedStock(dispatch)
	if !errors.Is(err, ErrInsufficientReservedStock) {
		t.Fatalf("expected ErrInsufficientReservedStock, got %v", err)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf("expected on hand 100, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40, got %d", inv.Reserved().Value())
	}

	if events := inv.Pull(); len(events) != 0 {
		t.Fatalf("expected no events on failed dispatch, got %d", len(events))
	}
}
