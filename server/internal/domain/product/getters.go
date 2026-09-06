package product

import "github.com/mwenza/mwenza/internal/platform/ids"

func (p *Product) ID() ids.ID {
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
