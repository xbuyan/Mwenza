package money

func (m Money) IsZero() bool {
	return m.amount == 0
}
