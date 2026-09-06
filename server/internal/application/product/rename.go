package product

import (
	"context"

	"github.com/mwenza/mwenza/internal/platform/ids"
)

type RenameCommand struct {
	ID   ids.ID
	Name string
}

type RenameHandler struct {
	repository Repository
}

func NewRenameHandler(repository Repository) *RenameHandler {
	return &RenameHandler{
		repository: repository,
	}
}

func (h *RenameHandler) Handle(
	ctx context.Context,
	cmd RenameCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return translateRepositoryError(err)
	}

	if product == nil {
		return ErrProductNotFound
	}

	if err := product.Rename(cmd.Name); err != nil {
		return err
	}

	return h.repository.Update(ctx, product)
}
