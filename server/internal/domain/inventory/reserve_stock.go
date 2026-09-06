package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrInvalidReservationQuantity = errors.New("reservation quantity must be greater than zero")
	ErrInsufficientAvailableStock = errors.New("insufficient available stock")
)

func (i *Inventory) ReserveStock(q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidReservationQuantity
	}

	available, err := i.Available()
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	remaining, err := available.Subtract(q)
	if err != nil {
		return ErrInsufficientAvailableStock
	}

	i.reserved = i.reserved.Add(q)

	i.Record(
		inventoryevents.NewStockReserved(
			i.productID,
			q,
		),
	)

	return nil
}
