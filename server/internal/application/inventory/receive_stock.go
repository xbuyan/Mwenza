package inventory

import (
	"context"
	"errors"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type ReceiveStockCommand struct {
	ProductID ids.ID
	Quantity  quantity.Quantity
}

type ReceiveStockHandler struct {
	repository Repository
}

func NewReceiveStockHandler(repository Repository) *ReceiveStockHandler {
	return &ReceiveStockHandler{
		repository: repository,
	}
}

func (h *ReceiveStockHandler) Handle(
	ctx context.Context,
	cmd ReceiveStockCommand,
) error {
	inventory, err := h.repository.FindByProductID(ctx, cmd.ProductID)

	if err != nil {
		if !errors.Is(err, ErrRepositoryNotFound) {
			return err
		}

		inventory, err = domaininventory.New(cmd.ProductID)
		if err != nil {
			return err
		}

		if err := inventory.ReceiveStock(cmd.Quantity); err != nil {
			return err
		}

		return h.repository.Save(ctx, inventory)
	}

	if err := inventory.ReceiveStock(cmd.Quantity); err != nil {
		return err
	}

	return h.repository.Update(ctx, inventory)
}
