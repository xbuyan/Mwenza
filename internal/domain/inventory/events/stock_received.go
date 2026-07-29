package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockReceivedEventName = "inventory.stock_received"

type StockReceived struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
}

func NewStockReceived(
	productID ids.ID,
	qty quantity.Quantity,
) StockReceived {
	return StockReceived{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
	}
}

func (e StockReceived) EventName() string {
	return StockReceivedEventName
}
