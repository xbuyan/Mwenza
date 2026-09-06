package inventory

import (
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReleaseReservedStock(t *testing.T) {
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

	// Clear events produced by receiving and reserving stock.
	inv.Pull()

	release, err := quantity.New(15)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReleaseReservedStock(release); err != nil {
		t.Fatal(err)
	}

	if inv.Reserved().Value() != 25 {
		t.Fatalf(
			"expected reserved 25, got %d",
			inv.Reserved().Value(),
		)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf(
			"expected on hand 100, got %d",
			inv.OnHand().Value(),
		)

	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 75 {
		t.Fatalf(
			"expected available 75, got %d",
			available.Value(),
		)
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockReservationReleased)
	if !ok {
		t.Fatalf(
			"expected StockReservationReleased event, got %T",
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

	if !event.Quantity.Equal(release) {
		t.Fatalf(
			"expected quantity %d, got %d",
			release.Value(),
			event.Quantity.Value(),
		)
	}
}

func TestReleaseReservedStockCannotExceedReserved(t *testing.T) {
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

	reserved, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	inv.Pull()

	release, err := quantity.New(6)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.ReleaseReservedStock(release)

	if err != ErrInsufficientReservedStock {
		t.Fatalf(
			"expected %v, got %v",
			ErrInsufficientReservedStock,
			err,
		)
	}

	if inv.Reserved().Value() != 5 {
		t.Fatalf(
			"reserved must remain 5, got %d",
			inv.Reserved().Value(),
		)
	}

	if inv.OnHand().Value() != 10 {
		t.Fatalf(
			"on hand must remain 10, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.HasEvents() {
		t.Fatal("failed release must not record an event")
	}
}

func TestReleaseReservedStockRejectsZeroQuantity(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.ReleaseReservedStock(zero)

	if err != ErrInvalidReleaseQuantity {
		t.Fatalf(
			"expected %v, got %v",
			ErrInvalidReleaseQuantity,
			err,
		)
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf(
			"reserved must remain 0, got %d",
			inv.Reserved().Value(),
		)
	}

	if inv.HasEvents() {
		t.Fatal("failed release must not record an event")
	}
}
