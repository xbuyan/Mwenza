package product

import "testing"

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
			id:      "prod-001",
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
			id:      "prod-001",
			sku:     "",
			product: "Bamburi Cement 50kg",
			unit:    UnitBag,
			wantErr: ErrEmptySKU,
		},
		{
			name:    "missing name",
			id:      "prod-001",
			sku:     "CEM-001",
			product: "",
			unit:    UnitBag,
			wantErr: ErrEmptyName,
		},
		{
			name:    "missing unit",
			id:      "prod-001",
			sku:     "CEM-001",
			product: "Bamburi Cement 50kg",
			unit:    "",
			wantErr: ErrInvalidUnit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.id, tt.sku, tt.product, tt.unit)

			if err != tt.wantErr {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
