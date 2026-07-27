package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDispatchQuantity = errors.New("dispatch quantity must be greater than zero")

func (i *Inventory) DispatchReservedStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidDispatchQuantity
	}

	reservedRemaining, err := i.reserved.Subtract(q)
	if err != nil {
		return ErrInsufficientReservedStock
	}

	onHandRemaining, err := i.onHand.Subtract(q)
	if err != nil {
		// This should never happen if inventory invariants are respected.
		return err
	}

	i.reserved = reservedRemaining
	i.onHand = onHandRemaining

	return nil
}
