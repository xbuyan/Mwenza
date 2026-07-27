package product

type Product struct {
	id     string
	sku    SKU
	name   string
	unit   Unit
	status Status
}

func New(id string, sku SKU, name string, unit Unit) (*Product, error) {
	if id == "" {
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
		id:     id,
		sku:    sku,
		name:   name,
		unit:   unit,
		status: StatusActive,
	}, nil
}
