package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestIsZero(t *testing.T) {
	tests := []struct {
		name   string
		money  Money
		expect bool
	}{
		{
			name:   "zero amount",
			money:  New(0, currency.KES),
			expect: true,
		},
		{
			name:   "positive amount",
			money:  New(100, currency.KES),
			expect: false,
		},
		{
			name:   "negative amount",
			money:  New(-100, currency.KES),
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.money.IsZero(); got != tt.expect {
				t.Fatalf("expected %v got %v", tt.expect, got)
			}
		})
	}
}
