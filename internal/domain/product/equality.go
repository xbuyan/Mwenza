package product

func (p *Product) Equals(other *Product) bool {
	if p == nil || other == nil {
		return false
	}

	return p.id == other.id
}
