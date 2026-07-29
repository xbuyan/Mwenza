package money

func (m Money) Equal(other Money) bool {
	return m.amount == other.amount &&
		m.currency == other.currency
}
