package money

func (m Money) Compare(other Money) (int, error) {
	if m.currency != other.currency {
		return 0, ErrCurrencyMismatch
	}

	switch {
	case m.amount < other.amount:
		return -1, nil
	case m.amount > other.amount:
		return 1, nil
	default:
		return 0, nil
	}
}
