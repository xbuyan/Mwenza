package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

func TestStockCount(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(100)
	_ = inv.ReceiveStock(start)

	counted, _ := quantity.New(92)

	if err := inv.StockCount(counted); err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 92 {
		t.Fatalf("expected on hand 92")
	}
}

func TestStockCountCannotGoBelowReserved(t *testing.T) {
	inv, _ := New("prod-001")

	start, _ := quantity.New(100)
	_ = inv.ReceiveStock(start)

	reserved, _ := quantity.New(30)
	_ = inv.ReserveStock(reserved)

	counted, _ := quantity.New(20)

	err := inv.StockCount(counted)

	if err != ErrCountBelowReserved {
		t.Fatalf("expected %v got %v", ErrCountBelowReserved, err)
	}
}
