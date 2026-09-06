package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type DiscontinueCommand struct {
	ID ids.ID
}

type DiscontinueHandler struct {
	repository Repository
}

func NewDiscontinueHandler(repository Repository) *DiscontinueHandler {
	return &DiscontinueHandler{
		repository: repository,
	}
}

func (h *DiscontinueHandler) Handle(
	ctx context.Context,
	cmd DiscontinueCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return translateRepositoryError(err)
	}

	if product == nil {
		return ErrProductNotFound
	}

	product.Discontinue()

	return h.repository.Update(ctx, product)
}
