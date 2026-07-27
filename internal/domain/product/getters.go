package product

func (p *Product) ID() string {
	return p.id
}

func (p *Product) SKU() SKU {
	return p.sku
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Unit() Unit {
	return p.unit
}

func (p *Product) Status() Status {
	return p.status
}
