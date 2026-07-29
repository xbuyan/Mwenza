package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestIsZero(t *testing.T) {
	if !New(0, currency.KES).IsZero() {
		t.Fatal("expected zero")
	}

	if New(1, currency.KES).IsZero() {
		t.Fatal("expected non-zero")
	}
}
