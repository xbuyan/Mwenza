package sale

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

func TestCreateSale(t *testing.T) {
	id := ids.New()

	s, err := New(id)
	if err != nil {
		t.Fatal(err)
	}

	if s.id != id {
		t.Fatalf("unexpected sale id")
	}

	if s.status != StatusDraft {
		t.Fatalf("expected draft status")
	}
}

func TestCreateSaleWithoutID(t *testing.T) {
	var id ids.ID

	_, err := New(id)

	if err != ErrEmptySaleID {
		t.Fatalf("expected %v got %v", ErrEmptySaleID, err)
	}
}
