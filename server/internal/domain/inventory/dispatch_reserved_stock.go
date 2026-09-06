package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDispatchQuantity = errors.New(
	"dispatch quantity must be greater than zero",
)

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

	i.Record(
		inventoryevents.NewStockDispatched(
			i.productID,
			q,
		),
	)

	return nil
}
