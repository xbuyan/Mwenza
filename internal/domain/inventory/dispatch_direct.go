package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidDirectDispatchQuantity = errors.New(
	"direct dispatch quantity must be greater than zero",
)

func (i *Inventory) DispatchDirect(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidDirectDispatchQuantity
	}

	available, err := i.Available()
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	if _, err := available.Subtract(q); err != nil {
		return ErrInsufficientAvailableStock
	}

	onHandRemaining, err := i.onHand.Subtract(q)
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	i.onHand = onHandRemaining

	i.Record(
		inventoryevents.NewStockDispatched(
			i.productID,
			q,
		),
	)

	return nil
}
