package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockCountedEventName = "inventory.stock_counted"

type StockCounted struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
}

func NewStockCounted(
	productID ids.ID,
	qty quantity.Quantity,
) StockCounted {
	return StockCounted{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
	}
}

func (e StockCounted) EventName() string {
	return StockCountedEventName
}
