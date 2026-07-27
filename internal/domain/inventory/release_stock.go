package inventory

import (
	"errors"

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

	return nil
}
