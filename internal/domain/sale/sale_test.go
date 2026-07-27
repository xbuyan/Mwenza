package sale

import "testing"

func TestSaleCanBeCreated(t *testing.T) {
	s := Sale{}

	if s.id != "" {
		t.Fatal("expected empty id")
	}
}
