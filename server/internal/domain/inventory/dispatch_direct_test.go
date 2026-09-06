package inventory

import (
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestDispatchDirect(t *testing.T) {
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

	dispatch, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	// Clear the stock.received event.
	inv.Pull()

	if err := inv.DispatchDirect(dispatch); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 70 {
		t.Fatalf("expected on hand 70, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected reserved 0, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 70 {
		t.Fatalf("expected available 70, got %d", available.Value())
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockDispatched)
	if !ok {
		t.Fatalf(
			"expected StockDispatched event, got %T",
			events[0],
		)
	}

	if event.ProductID != inv.ProductID() {
		t.Fatalf(
			"expected product ID %v, got %v",
			inv.ProductID(),
			event.ProductID,
		)
	}

	if !event.Quantity.Equal(dispatch) {
		t.Fatalf(
			"expected quantity %d, got %d",
			dispatch.Value(),
			event.Quantity.Value(),
		)
	}
}

func TestDispatchDirectInsufficientStock(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	inv.Pull()

	dispatch, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.DispatchDirect(dispatch)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf(
			"expected %v, got %v",
			ErrInsufficientAvailableStock,
			err,
		)
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf(
			"expected on hand 20 after failed dispatch, got %d",
			inv.OnHand().Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 0 {
		t.Fatalf(
			"expected 0 events after failed dispatch, got %d",
			len(events),
		)
	}
}

func TestDispatchDirectZeroQuantity(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.DispatchDirect(zero)

	if err != ErrInvalidDirectDispatchQuantity {
		t.Fatalf(
			"expected %v, got %v",
			ErrInvalidDirectDispatchQuantity,
			err,
		)
	}

	events := inv.Pull()

	if len(events) != 0 {
		t.Fatalf(
			"expected 0 events after failed dispatch, got %d",
			len(events),
		)
	}
}

func TestDispatchDirectCannotConsumeReservedStock(t *testing.T) {
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

	dispatch, err := quantity.New(70)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.DispatchDirect(dispatch)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf(
			"expected %v, got %v",
			ErrInsufficientAvailableStock,
			err,
		)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf(
			"expected on hand 100 after failed dispatch, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.Reserved().Value() != 40 {
		t.Fatalf(
			"expected reserved 40 after failed dispatch, got %d",
			inv.Reserved().Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 0 {
		t.Fatalf(
			"expected 0 events after failed dispatch, got %d",
			len(events),
		)
	}
}
