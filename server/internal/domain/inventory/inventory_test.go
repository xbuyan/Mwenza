package inventory

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

func TestCreateInventory(t *testing.T) {
	inv, err := New(ids.New())
	if err != nil {
		t.Fatal(err)
	}

	if inv.OnHand().Value() != 0 {
		t.Fatalf("expected zero on hand")
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected zero reserved")
	}
}

func TestCreateInventoryWithoutProduct(t *testing.T) {
	var id ids.ID

	_, err := New(id)

	if err != ErrEmptyProductID {
		t.Fatalf("expected %v got %v", ErrEmptyProductID, err)
	}
}
