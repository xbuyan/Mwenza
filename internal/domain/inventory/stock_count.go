package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrCountBelowReserved = errors.New("counted stock cannot be below reserved stock")

func (i *Inventory) StockCount(actual quantity.Quantity) error {
	if _, err := actual.Subtract(i.reserved); err != nil {
		return ErrCountBelowReserved
	}

	i.onHand = actual
	return nil
}
