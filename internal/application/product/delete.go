package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type DeleteCommand struct {
	ID ids.ID
}

type DeleteHandler struct {
	repository Repository
}

func NewDeleteHandler(repository Repository) *DeleteHandler {
	return &DeleteHandler{
		repository: repository,
	}
}

func (h *DeleteHandler) Handle(
	ctx context.Context,
	cmd DeleteCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	return h.repository.Delete(ctx, cmd.ID)
}
