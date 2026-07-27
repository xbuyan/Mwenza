package quantity

import "testing"

func TestNegativeQuantity(t *testing.T) {
	_, err := New(-1)

	if err == nil {
		t.Fatal("expected error")
	}
}
