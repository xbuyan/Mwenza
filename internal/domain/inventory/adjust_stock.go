package inventory

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type AdjustmentDirection int

const (
	Increase AdjustmentDirection = iota
	Decrease
)

var (
	ErrInvalidAdjustmentQuantity = errors.New("adjustment quantity must be greater than zero")
	ErrAdjustmentBelowReserved   = errors.New("adjustment would reduce stock below reserved quantity")
)

func (i *Inventory) AdjustStock(direction AdjustmentDirection, q quantity.Quantity) error {
	if q.IsZero() {
		return ErrInvalidAdjustmentQuantity
	}

	switch direction {
	case Increase:
		i.onHand = i.onHand.Add(q)
		return nil

	case Decrease:
		newOnHand, err := i.onHand.Subtract(q)
		if err != nil {
			return ErrAdjustmentBelowReserved
		}

		if _, err := newOnHand.Subtract(i.reserved); err != nil {
			return ErrAdjustmentBelowReserved
		}

		i.onHand = newOnHand
		return nil

	default:
		return errors.New("invalid adjustment direction")
	}
}
