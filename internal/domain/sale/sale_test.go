package sale

import "testing"

func TestCreateSale(t *testing.T) {
	s, err := New("sale-001")
	if err != nil {
		t.Fatal(err)
	}

	if s.id != "sale-001" {
		t.Fatalf("unexpected sale id")
	}

	if s.status != StatusDraft {
		t.Fatalf("expected draft status")
	}
}

func TestCreateSaleWithoutID(t *testing.T) {
	_, err := New("")

	if err != ErrEmptySaleID {
		t.Fatalf("expected %v got %v", ErrEmptySaleID, err)
	}
}
