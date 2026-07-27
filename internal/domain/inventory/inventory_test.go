package inventory

import "testing"

func TestInventoryCanBeCreated(t *testing.T) {
	i := Inventory{}

	if i.onHand != 0 {
		t.Fatalf("expected zero stock")
	}
}
