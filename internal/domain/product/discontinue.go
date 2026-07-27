package product

func (p *Product) Discontinue() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	p.status = StatusDiscontinued
}
