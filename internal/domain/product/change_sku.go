package product

func (p *Product) ChangeSKU(sku SKU) error {
	if sku == "" {
		return ErrEmptySKU
	}

	if p.sku == sku {
		return nil
	}

	p.sku = sku
	return nil
}
