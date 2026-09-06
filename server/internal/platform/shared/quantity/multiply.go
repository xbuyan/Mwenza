package quantity

func (q Quantity) Multiply(factor int64) (Quantity, error) {
	if factor < 0 {
		return Quantity{}, ErrNegativeQuantity
	}

	return Quantity{
		value: q.value * factor,
	}, nil
}
