package inventory

import (
	"testing"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestIncreaseStock(t *testing.T) {
	inv, _ := New("prod-001")

	q, _ := quantity.New(20)

	if err := inv.AdjustStock(Increase, q); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf("expected on hand 20, got %d", inv.OnHand().Value())
	}
}

func TestDecreaseStock(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(50)
	_ = inv.ReceiveStock(start)

	dec, _ := quantity.New(10)

	if err := inv.AdjustStock(Decrease, dec); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 40 {
		t.Fatalf("expected on hand 40, got %d", inv.OnHand().Value())
	}
}

func TestCannotDecreaseBelowZero(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(5)
	_ = inv.ReceiveStock(start)

	dec, _ := quantity.New(10)

	err := inv.AdjustStock(Decrease, dec)

	if err != ErrInsufficientStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientStock, err)
	}

	if inv.OnHand().Value() != 5 {
		t.Fatalf("expected on hand to remain 5, got %d", inv.OnHand().Value())
	}
}

func TestCannotAdjustBelowReserved(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(20)
	_ = inv.ReceiveStock(start)

	reserved, _ := quantity.New(15)
	_ = inv.ReserveStock(reserved)

	dec, _ := quantity.New(10)

	err := inv.AdjustStock(Decrease, dec)

	if err != ErrAdjustmentBelowReserved {
		t.Fatalf("expected %v, got %v", ErrAdjustmentBelowReserved, err)
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf("expected on hand to remain 20, got %d", inv.OnHand().Value())
	}
}

func TestDecreaseExactlyToReservedQuantity(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(20)
	_ = inv.ReceiveStock(start)

	reserved, _ := quantity.New(15)
	_ = inv.ReserveStock(reserved)

	dec, _ := quantity.New(5)

	if err := inv.AdjustStock(Decrease, dec); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 15 {
		t.Fatalf("expected on hand 15, got %d", inv.OnHand().Value())
	}
}

func TestRejectsZeroAdjustment(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.AdjustStock(Increase, zero)

	if err != ErrInvalidAdjustmentQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidAdjustmentQuantity, err)
	}
}

func TestRejectsInvalidAdjustmentDirection(t *testing.T) {
	inv, _ := New("prod-001")

	q, _ := quantity.New(10)

	err := inv.AdjustStock(AdjustmentDirection(99), q)

	if err != ErrInvalidAdjustmentDirection {
		t.Fatalf("expected %v, got %v", ErrInvalidAdjustmentDirection, err)
	}
}

func TestInvalidAdjustmentDoesNotChangeStock(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(50)
	_ = inv.ReceiveStock(start)

	q, _ := quantity.New(10)

	err := inv.AdjustStock(AdjustmentDirection(99), q)

	if err != ErrInvalidAdjustmentDirection {
		t.Fatalf("expected %v, got %v", ErrInvalidAdjustmentDirection, err)
	}

	if inv.OnHand().Value() != 50 {
		t.Fatalf("expected on hand to remain 50, got %d", inv.OnHand().Value())
	}
}

func TestIncreaseStockRecordsEvent(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	q, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.AdjustStock(Increase, q); err != nil {
		t.Fatal(err)
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockAdjusted)
	if !ok {
		t.Fatalf("expected StockAdjusted event, got %T", events[0])
	}

	if event.ProductID != inv.ProductID() {
		t.Fatalf(
			"expected product ID %v, got %v",
			inv.ProductID(),
			event.ProductID,
		)
	}

	if !event.Quantity.Equal(q) {
		t.Fatalf(
			"expected quantity %d, got %d",
			q.Value(),
			event.Quantity.Value(),
		)
	}

	if event.Direction != int(Increase) {
		t.Fatalf(
			"expected direction %d, got %d",
			int(Increase),
			event.Direction,
		)
	}
}

func TestDecreaseStockRecordsEvent(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	start, err := quantity.New(50)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(start); err != nil {
		t.Fatal(err)
	}

	// Remove the stock.received event.
	inv.Pull()

	dec, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.AdjustStock(Decrease, dec); err != nil {
		t.Fatal(err)
	}

	events := inv.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockAdjusted)
	if !ok {
		t.Fatalf("expected StockAdjusted event, got %T", events[0])
	}

	if event.ProductID != inv.ProductID() {
		t.Fatalf(
			"expected product ID %v, got %v",
			inv.ProductID(),
			event.ProductID,
		)
	}

	if !event.Quantity.Equal(dec) {
		t.Fatalf(
			"expected quantity %d, got %d",
			dec.Value(),
			event.Quantity.Value(),
		)
	}

	if event.Direction != int(Decrease) {
		t.Fatalf(
			"expected direction %d, got %d",
			int(Decrease),
			event.Direction,
		)
	}
}
func TestInsufficientStockRecordsNoEvent(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	start, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(start); err != nil {
		t.Fatal(err)
	}

	// Remove the stock.received event.
	inv.Pull()

	dec, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.AdjustStock(Decrease, dec)

	if err != ErrInsufficientStock {
		t.Fatalf(
			"expected %v, got %v",
			ErrInsufficientStock,
			err,
		)
	}

	if inv.HasEvents() {
		t.Fatal("failed adjustment must not record an event")
	}

	if inv.OnHand().Value() != 5 {
		t.Fatalf(
			"expected on hand 5, got %d",
			inv.OnHand().Value(),
		)
	}
}

func TestAdjustmentBelowReservedRecordsNoEvent(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	start, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(start); err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(15)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	// Remove the events produced by receiving and reserving stock.
	inv.Pull()

	dec, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = inv.AdjustStock(Decrease, dec)

	if err != ErrAdjustmentBelowReserved {
		t.Fatalf(
			"expected %v, got %v",
			ErrAdjustmentBelowReserved,
			err,
		)
	}

	if inv.HasEvents() {
		t.Fatal("failed adjustment must not record an event")
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf(
			"expected on hand 20, got %d",
			inv.OnHand().Value(),
		)
	}

	if inv.Reserved().Value() != 15 {
		t.Fatalf(
			"expected reserved 15, got %d",
			inv.Reserved().Value(),
		)
	}
}
