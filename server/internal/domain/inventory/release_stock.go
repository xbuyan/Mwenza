package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrInvalidReleaseQuantity    = errors.New("release quantity must be greater than zero")
	ErrInsufficientReservedStock = errors.New("insufficient reserved stock")
)

func (i *Inventory) ReleaseReservedStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReleaseQuantity
	}

	remaining, err := i.reserved.Subtract(q)
	if err != nil {
		return ErrInsufficientReservedStock
	}

	i.reserved = remaining

	i.Record(
		inventoryevents.NewStockReservationReleased(
			i.productID,
			q,
		),
	)

	return nil
}
