package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrInvalidReceiveQuantity = errors.New("received quantity must be greater than zero")

func (i *Inventory) ReceiveStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReceiveQuantity
	}

	i.onHand = i.onHand.Add(q)

	i.Record(
		inventoryevents.NewStockReceived(
			i.productID,
			q,
		),
	)

	return nil
}
