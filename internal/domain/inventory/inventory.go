package inventory

import "errors"

var ErrEmptyProductID = errors.New("product id cannot be empty")

type Inventory struct {
	productID string
	onHand    int
	reserved  int
}

func New(productID string) (*Inventory, error) {
	if productID == "" {
		return nil, ErrEmptyProductID
	}

	return &Inventory{
		productID: productID,
		onHand:    0,
		reserved:  0,
	}, nil
}
