package product

func (p *Product) Description() string {
	return p.description
}

func (p *Product) ChangeDescription(description string) {
	if p == nil {
		return
	}

	p.description = description
}
