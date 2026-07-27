package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidReceiveQuantity = errors.New("received quantity must be greater than zero")

func (i *Inventory) ReceiveStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReceiveQuantity
	}

	i.onHand = i.onHand.Add(q)

	return nil
}
