package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockReservedEventName = "inventory.stock_reserved"

type StockReserved struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
}

func NewStockReserved(
	productID ids.ID,
	qty quantity.Quantity,
) StockReserved {
	return StockReserved{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
	}
}

func (e StockReserved) EventName() string {
	return StockReservedEventName
}
