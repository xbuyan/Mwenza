package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrCountBelowReserved = errors.New("counted stock cannot be below reserved stock")

func (i *Inventory) StockCount(actual quantity.Quantity) error {
	if _, err := actual.Subtract(i.reserved); err != nil {
		return ErrCountBelowReserved
	}

	i.onHand = actual

	i.Record(
		inventoryevents.NewStockCounted(
			i.productID,
			actual,
		),
	)

	return nil
}
