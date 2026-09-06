package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockDispatchedEventName = "inventory.stock_dispatched"

type StockDispatched struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
}

func NewStockDispatched(
	productID ids.ID,
	qty quantity.Quantity,
) StockDispatched {
	return StockDispatched{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
	}
}

func (e StockDispatched) EventName() string {
	return StockDispatchedEventName
}
