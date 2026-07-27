package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestReleaseReservedStock(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(100)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(40)
	_ = inv.ReserveStock(reserved)

	release, _ := quantity.New(15)

	if err := inv.ReleaseReservedStock(release); err != nil {
		t.Fatal(err)
	}

	if inv.Reserved().Value() != 25 {
		t.Fatalf("expected reserved 25, got %d", inv.Reserved().Value())
	}

	available, _ := inv.Available()

	if available.Value() != 75 {
		t.Fatalf("expected available 75, got %d", available.Value())
	}
}

func TestReleaseTooMuch(t *testing.T) {
	inv, _ := New("prod-001")

	received, _ := quantity.New(10)
	_ = inv.ReceiveStock(received)

	reserved, _ := quantity.New(5)
	_ = inv.ReserveStock(reserved)

	release, _ := quantity.New(6)

	err := inv.ReleaseReservedStock(release)

	if err != ErrInsufficientReservedStock {
		t.Fatalf("expected %v got %v", ErrInsufficientReservedStock, err)
	}
}

func TestReleaseZero(t *testing.T) {
	inv, _ := New("prod-001")

	zero, _ := quantity.New(0)

	err := inv.ReleaseReservedStock(zero)

	if err != ErrInvalidReleaseQuantity {
		t.Fatalf("expected %v got %v", ErrInvalidReleaseQuantity, err)
	}
}
