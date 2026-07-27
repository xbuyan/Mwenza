package money

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

var ErrCurrencyMismatch = errors.New("currency mismatch")

type Money struct {
	amount   int64
	currency currency.Currency
}

func New(amount int64, curr currency.Currency) Money {
	return Money{
		amount:   amount,
		currency: curr,
	}
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() currency.Currency {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, ErrCurrencyMismatch
	}

	return Money{
		amount:   m.amount + other.amount,
		currency: m.currency,
	}, nil
}
