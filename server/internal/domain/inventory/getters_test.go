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

	if inv.OnHand().Value() != 0 {
		t.Fatalf("expected onHand = 0")
	}

	if inv.Reserved().Value() != 0 {
		t.Fatalf("expected reserved = 0")
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 0 {
		t.Fatalf("expected available = 0")
	}
}
