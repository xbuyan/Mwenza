package money

func (m Money) Multiply(factor int64) Money {
	return Money{
		amount:   m.amount * factor,
		currency: m.currency,
	}
}
