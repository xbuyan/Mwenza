package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestIncreaseStock(t *testing.T) {
	inv, _ := New("prod-001")

	q, _ := quantity.New(20)

	if err := inv.AdjustStock(Increase, q); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 20 {
		t.Fatalf("expected on hand 20")
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
		t.Fatalf("expected on hand 40")
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
		t.Fatalf("expected %v got %v", ErrAdjustmentBelowReserved, err)
	}
}
