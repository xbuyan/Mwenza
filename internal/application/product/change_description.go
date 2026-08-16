package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type ChangeDescriptionCommand struct {
	ID          ids.ID
	Description string
}

type ChangeDescriptionHandler struct {
	repository Repository
}

func NewChangeDescriptionHandler(repository Repository) *ChangeDescriptionHandler {
	return &ChangeDescriptionHandler{
		repository: repository,
	}
}

func (h *ChangeDescriptionHandler) Handle(
	ctx context.Context,
	cmd ChangeDescriptionCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	product.ChangeDescription(cmd.Description)

	return h.repository.Update(ctx, product)
}
