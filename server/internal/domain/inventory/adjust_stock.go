package inventory

import (
	"errors"

	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type AdjustmentDirection int

const (
	Increase AdjustmentDirection = iota
	Decrease
)

var (
	ErrInvalidAdjustmentQuantity  = errors.New("adjustment quantity must be greater than zero")
	ErrInvalidAdjustmentDirection = errors.New("invalid adjustment direction")
	ErrInsufficientStock          = errors.New("insufficient stock")
	ErrAdjustmentBelowReserved    = errors.New("adjustment would reduce stock below reserved quantity")
)

func (i *Inventory) AdjustStock(direction AdjustmentDirection, q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidAdjustmentQuantity
	}

	switch direction {
	case Increase:
		i.onHand = i.onHand.Add(q)

	case Decrease:
		newOnHand, err := i.onHand.Subtract(q)
		if err != nil {
			return ErrInsufficientStock
		}

		if _, err := newOnHand.Subtract(i.reserved); err != nil {
			return ErrAdjustmentBelowReserved
		}

		i.onHand = newOnHand

	default:
		return ErrInvalidAdjustmentDirection
	}

	i.Record(
		inventoryevents.NewStockAdjusted(
			i.productID,
			q,
			int(direction),
		),
	)

	return nil
}
