package inventory

import (
	"context"
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type StockCountCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type StockCountHandler struct {
	repository Repository
}

func NewStockCountHandler(repository Repository) *StockCountHandler {
	return &StockCountHandler{
		repository: repository,
	}
}

func (h *StockCountHandler) Handle(
	ctx context.Context,
	cmd StockCountCommand,
) error {
	inventory, err := h.repository.FindByProductID(ctx, cmd.ProductID)
	if err != nil {
		if errors.Is(err, ErrRepositoryNotFound) {
			return ErrInventoryNotFound
		}

		return err
	}

	if inventory == nil {
		return ErrInventoryNotFound
	}

	if err := inventory.StockCount(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
