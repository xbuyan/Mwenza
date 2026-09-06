package inventory

import (
	"errors"

	platformevents "github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var ErrEmptyProductID = errors.New("product id cannot be empty")

type Inventory struct {
	platformevents.AggregateRoot

	productID ids.ID
	onHand    quantity.Quantity
	reserved  quantity.Quantity
}

func New(productID ids.ID) (*Inventory, error) {
	if productID.IsZero() {
		return nil, ErrEmptyProductID
	}

	zero, err := quantity.New(0)
	if err != nil {
		return nil, err
	}

	return &Inventory{
		AggregateRoot: platformevents.AggregateRoot{},
		productID:     productID,
		onHand:        zero,
		reserved:      zero,
	}, nil
}
