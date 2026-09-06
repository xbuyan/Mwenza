package inventory

import (
	"context"
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type DispatchDirectStockCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type DispatchDirectStockHandler struct {
	repository Repository
}

func NewDispatchDirectStockHandler(repository Repository) *DispatchDirectStockHandler {
	return &DispatchDirectStockHandler{
		repository: repository,
	}
}

func (h *DispatchDirectStockHandler) Handle(
	ctx context.Context,
	cmd DispatchDirectStockCommand,
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

	if err := inventory.DispatchDirect(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
