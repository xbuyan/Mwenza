package product

func (p *Product) ChangeUnit(unit Unit) error {
	if unit == "" {
		return ErrInvalidUnit
	}

	if p.unit == unit {
		return nil
	}

	p.unit = unit
	return nil
}
