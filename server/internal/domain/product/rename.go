package product

func (p *Product) Rename(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	p.name = name
	return nil
}
