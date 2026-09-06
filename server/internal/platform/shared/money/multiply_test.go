package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		factor   int64
		expected int64
	}{
		{
			name:     "positive multiplier",
			money:    New(250, currency.KES),
			factor:   4,
			expected: 1000,
		},
		{
			name:     "zero multiplier",
			money:    New(250, currency.KES),
			factor:   0,
			expected: 0,
		},
		{
			name:     "negative multiplier",
			money:    New(250, currency.KES),
			factor:   -2,
			expected: -500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.money.Multiply(tt.factor)

			if got.Amount() != tt.expected {
				t.Fatalf("expected %d got %d", tt.expected, got.Amount())
			}

			if got.Currency() != tt.money.Currency() {
				t.Fatal("currency changed")
			}
		})
	}
}
