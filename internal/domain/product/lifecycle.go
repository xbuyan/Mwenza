package product

func (p *Product) Activate() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	if p.status == StatusInactive {
		p.status = StatusActive
	}
}

func (p *Product) Deactivate() {
	if p == nil {
		return
	}

	if p.status == StatusDiscontinued {
		return
	}

	if p.status == StatusActive {
		p.status = StatusInactive
	}
}
