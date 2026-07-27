package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReserveStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserve, _ := quantity.New(30)

	if err := inv.ReserveStock(reserve); err != nil {
		t.Fatal(err)
	}

	if inv.Reserved().Value() != 30 {
		t.Fatalf("expected reserved 30, got %d", inv.Reserved().Value())
	}

	available, _ := inv.Available()

	if available.Value() != 70 {
		t.Fatalf("expected available 70, got %d", available.Value())
	}
}

func TestReserveTooMuchStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(10)
	_ = inv.ReceiveStock(received)

	reserve, _ := quantity.New(20)

	err := inv.ReserveStock(reserve)

	if err != ErrInsufficientAvailableStock {
		t.Fatalf("expected %v, got %v", ErrInsufficientAvailableStock, err)
	}
}

func TestReserveZeroStock(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.ReserveStock(zero)

	if err != ErrInvalidReservationQuantity {
		t.Fatalf("expected %v, got %v", ErrInvalidReservationQuantity, err)
	}
}
