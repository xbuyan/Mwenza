package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type ActivateCommand struct {
	ID ids.ID
}

type ActivateHandler struct {
	repository Repository
}

func NewActivateHandler(repository Repository) *ActivateHandler {
	return &ActivateHandler{
		repository: repository,
	}
}

func (h *ActivateHandler) Handle(
	ctx context.Context,
	cmd ActivateCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	product.Activate()

	return h.repository.Update(ctx, product)
}
