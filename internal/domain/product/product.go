package product

import "github.com/mwenza/mwenza/internal/platform/ids"

type Product struct {
	id          ids.ID
	sku         SKU
	name        string
	description string
	unit        Unit
	status      Status
}

func New(id ids.ID, sku SKU, name string, unit Unit) (*Product, error) {
	if id.IsZero() {
		return nil, ErrEmptyID
	}

	if sku == "" {
		return nil, ErrEmptySKU
	}

	if name == "" {
		return nil, ErrEmptyName
	}

	if unit == "" {
		return nil, ErrInvalidUnit
	}

	return &Product{
		id:          id,
		sku:         sku,
		name:        name,
		description: "",
		unit:        unit,
		status:      StatusActive,
	}, nil
}
