package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDirectDispatchQuantity = errors.New("direct dispatch quantity must be greater than zero")

func (i *Inventory) DispatchDirect(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidDirectDispatchQuantity
	}

	available, err := i.Available()
	if err != nil {
		return err
	}

	remainingAvailable, err := available.Subtract(q)
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	_ = remainingAvailable

	onHandRemaining, err := i.onHand.Subtract(q)
	if err != nil {
		return err
	}

	i.onHand = onHandRemaining

	return nil
}
