package inventory

import "testing"

func TestCreateInventory(t *testing.T) {
	inv, err := New("prod-001")

	if err != nil {
		t.Fatal(err)
	}

	if inv.onHand != 0 {
		t.Fatalf("expected zero on hand")
	}

	if inv.reserved != 0 {
		t.Fatalf("expected zero reserved")
	}
}

func TestCreateInventoryWithoutProduct(t *testing.T) {
	_, err := New("")

	if err != ErrEmptyProductID {
		t.Fatalf("expected %v got %v", ErrEmptyProductID, err)
	}
}
