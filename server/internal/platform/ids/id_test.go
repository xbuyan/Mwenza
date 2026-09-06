package ids

import "testing"

func TestNewID(t *testing.T) {
	id := New()

	if id.IsZero() {
		t.Fatal("expected non-zero id")
	}
}

func TestParseValidID(t *testing.T) {
	id := New()

	parsed, err := Parse(id.String())
	if err != nil {
		t.Fatal(err)
	}

	if parsed != id {
		t.Fatal("parsed id mismatch")
	}
}

func TestParseInvalidID(t *testing.T) {
	_, err := Parse("abc")

	if err != ErrInvalidID {
		t.Fatalf("expected %v got %v", ErrInvalidID, err)
	}
}

func TestMustParse(t *testing.T) {
	id := New()

	if MustParse(id.String()) != id {
		t.Fatal("unexpected id")
	}
}

func TestIsZero(t *testing.T) {
	var id ID

	if !id.IsZero() {
		t.Fatal("expected zero id")
	}
}
