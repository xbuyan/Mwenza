package inventory

import (
	"context"
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type DispatchReservedStockCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type DispatchReservedStockHandler struct {
	repository Repository
}

func NewDispatchReservedStockHandler(repository Repository) *DispatchReservedStockHandler {
	return &DispatchReservedStockHandler{
		repository: repository,
	}
}

func (h *DispatchReservedStockHandler) Handle(
	ctx context.Context,
	cmd DispatchReservedStockCommand,
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

	if err := inventory.DispatchReservedStock(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
