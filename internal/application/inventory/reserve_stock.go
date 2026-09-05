package inventory

import (
	"context"
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type ReserveStockCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type ReserveStockHandler struct {
	repository Repository
}

func NewReserveStockHandler(repository Repository) *ReserveStockHandler {
	return &ReserveStockHandler{
		repository: repository,
	}
}

func (h *ReserveStockHandler) Handle(
	ctx context.Context,
	cmd ReserveStockCommand,
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

	if err := inventory.ReserveStock(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
