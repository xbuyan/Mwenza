package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockAdjustedEventName = "inventory.stock_adjusted"

type StockAdjusted struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
	Direction int
}

func NewStockAdjusted(
	productID ids.ID,
	qty quantity.Quantity,
	direction int,
) StockAdjusted {
	return StockAdjusted{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
		Direction: direction,
	}
}

func (e StockAdjusted) EventName() string {
	return StockAdjustedEventName
}
