package inventory

import (
	"context"
	"errors"

	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type ReleaseReservedStockCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type ReleaseReservedStockHandler struct {
	repository Repository
}

func NewReleaseReservedStockHandler(repository Repository) *ReleaseReservedStockHandler {
	return &ReleaseReservedStockHandler{
		repository: repository,
	}
}

func (h *ReleaseReservedStockHandler) Handle(
	ctx context.Context,
	cmd ReleaseReservedStockCommand,
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

	if err := inventory.ReleaseReservedStock(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
