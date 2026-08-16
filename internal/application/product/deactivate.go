package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type DeactivateCommand struct {
	ID ids.ID
}

type DeactivateHandler struct {
	repository Repository
}

func NewDeactivateHandler(repository Repository) *DeactivateHandler {
	return &DeactivateHandler{
		repository: repository,
	}
}

func (h *DeactivateHandler) Handle(
	ctx context.Context,
	cmd DeactivateCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	product.Deactivate()

	return h.repository.Update(ctx, product)
}
