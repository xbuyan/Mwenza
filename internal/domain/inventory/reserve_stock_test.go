package inventory

import (
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReserveStock(t *testing.T) {
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

	reserve, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	inv.Pull()

	if err := inv.ReserveStock(reserve); err != nil {
		t.Fatal(err)
	}

	if inv.Reserved().Value() != 30 {
		t.Fatalf(
			"expected reserved 30, got %d",
			inv.Reserved().Value(),
		)
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 70 {
		t.Fatalf(
			"expected available 70, got %d",
			available.Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockReserved)
	if !ok {
		t.Fatalf(
			"expected StockReserved event, got %T",
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

	if !event.Quantity.Equal(reserve) {
		t.Fatalf(
			"expected quantity %d, got %d",
			reserve.Value(),
			event.Quantity.Value(),
		)
	}
}

func TestReserveTooMuchStock(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	inv.Pull()

	reserve, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.ReserveStock(reserve)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf(
			"expected %v, got %v",
			ErrInsufficientAvailableStock,
			err,
		)
	}

	if inv.OnHand().Value() != 10 {
		t.Fatalf(
			"expected on hand to remain 10, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf(
			"expected reserved to remain 0, got %d",
			inv.Reserved().Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 0 {
		t.Fatalf(
			"expected no event after failed reservation, got %d",
			len(events),
		)
	}
}

func TestReserveZeroStock(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.ReserveStock(zero)

	if err != ErrInvalidReservationQuantity {
		t.Fatalf(
			"expected %v, got %v",
			ErrInvalidReservationQuantity,
			err,
		)
	}

	if inv.OnHand().Value() != 0 {
		t.Fatalf(
			"expected on hand to remain 0, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf(
			"expected reserved to remain 0, got %d",
			inv.Reserved().Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 0 {
		t.Fatalf(
			"expected no event after failed reservation, got %d",
			len(events),
		)
	}
}
