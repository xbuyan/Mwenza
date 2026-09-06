package quantity

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		factor   int64
		expected int64
		wantErr  error
	}{
		{
			name:     "positive multiplier",
			value:    5,
			factor:   4,
			expected: 20,
		},
		{
			name:     "zero multiplier",
			value:    5,
			factor:   0,
			expected: 0,
		},
		{
			name:    "negative multiplier",
			value:   5,
			factor:  -2,
			wantErr: ErrNegativeQuantity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, _ := New(tt.value)

			got, err := q.Multiply(tt.factor)

			if err != tt.wantErr {
				t.Fatalf("expected %v got %v", tt.wantErr, err)
			}

			if err == nil && got.Value() != tt.expected {
				t.Fatalf("expected %d got %d", tt.expected, got.Value())
			}
		})
	}
}
