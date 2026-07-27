package sale

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/money"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

var (
	ErrEmptyProductID   = errors.New("product id cannot be empty")
	ErrEmptyProductName = errors.New("product name cannot be empty")
)

type LineItem struct {
	productID   string
	productName string
	quantity    quantity.Quantity
	unitPrice   money.Money
}

func NewLineItem(
	productID string,
	productName string,
	qty quantity.Quantity,
	price money.Money,
) (LineItem, error) {

	if productID == "" {
		return LineItem{}, ErrEmptyProductID
	}

	if productName == "" {
		return LineItem{}, ErrEmptyProductName
	}

	return LineItem{
		productID:   productID,
		productName: productName,
		quantity:    qty,
		unitPrice:   price,
	}, nil
}
