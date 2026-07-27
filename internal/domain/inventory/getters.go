package inventory

import "github.com/mwenza/mwenza/internal/platform/shared/quantity"

func (i *Inventory) ProductID() string {
	return i.productID
}

func (i *Inventory) OnHand() quantity.Quantity {
	return i.onHand
}

func (i *Inventory) Reserved() quantity.Quantity {
	return i.reserved
}

func (i *Inventory) Available() (quantity.Quantity, error) {
	return i.onHand.Subtract(i.reserved)
}
