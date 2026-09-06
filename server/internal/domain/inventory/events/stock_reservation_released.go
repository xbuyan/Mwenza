package events

import (
	"github.com/mwenza/mwenza/internal/platform/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

const StockReservationReleasedEventName = "inventory.stock_reservation_released"

type StockReservationReleased struct {
	events.BaseEvent

	ProductID ids.ID
	Quantity  quantity.Quantity
}

func NewStockReservationReleased(
	productID ids.ID,
	qty quantity.Quantity,
) StockReservationReleased {
	return StockReservationReleased{
		BaseEvent: events.NewBaseEvent(),
		ProductID: productID,
		Quantity:  qty,
	}
}

func (e StockReservationReleased) EventName() string {
	return StockReservationReleasedEventName
}
