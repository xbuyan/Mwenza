package inventory

import "testing"

func TestInventoryGetters(t *testing.T) {
	inv, err := New("prod-001")
	if err != nil {
		t.Fatal(err)
	}

	if inv.ProductID() != "prod-001" {
		t.Fatalf("unexpected product id")
	}

	if inv.OnHand() != 0 {
		t.Fatalf("expected onHand = 0")
	}

	if inv.Reserved() != 0 {
		t.Fatalf("expected reserved = 0")
	}

	if inv.Available() != 0 {
		t.Fatalf("expected available = 0")
	}
}
