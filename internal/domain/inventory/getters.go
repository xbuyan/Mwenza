package inventory

func (i *Inventory) ProductID() string {
	return i.productID
}

func (i *Inventory) OnHand() int {
	return i.onHand
}

func (i *Inventory) Reserved() int {
	return i.reserved
}

func (i *Inventory) Available() int {
	return i.onHand - i.reserved
}
