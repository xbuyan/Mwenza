package product

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		sku     SKU
		product string
		unit    Unit
		wantErr error
	}{
		{
			name:    "valid product",
			id:      ids.New().String(),
			sku:     "CEM-001",
			product: "Bamburi Cement 50kg",
			unit:    UnitBag,
			wantErr: nil,
		},
		{
			name:    "missing id",
			id:      "",
			sku:     "CEM-001",
			product: "Bamburi Cement 50kg",
			unit:    UnitBag,
			wantErr: ErrEmptyID,
		},
		{
			name:    "missing sku",
			id:      ids.New().String(),
			sku:     "",
			product: "Bamburi Cement 50kg",
			unit:    UnitBag,
			wantErr: ErrEmptySKU,
		},
		{
			name:    "missing name",
			id:      ids.New().String(),
			sku:     "CEM-001",
			product: "",
			unit:    UnitBag,
			wantErr: ErrEmptyName,
		},
		{
			name:    "missing unit",
			id:      ids.New().String(),
			sku:     "CEM-001",
			product: "Bamburi Cement 50kg",
			unit:    "",
			wantErr: ErrInvalidUnit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var id ids.ID

			if tt.id != "" {
				var err error
				id, err = ids.Parse(tt.id)
				if err != nil {
					t.Fatalf("unexpected test id: %v", err)
				}
			}

			_, err := New(id, tt.sku, tt.product, tt.unit)

			if err != tt.wantErr {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
