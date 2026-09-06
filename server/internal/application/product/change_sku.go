package product

import (
	"context"
	"errors"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type ChangeSKUCommand struct {
	ID  ids.ID
	SKU domainproduct.SKU
}

type ChangeSKUHandler struct {
	repository Repository
}

func NewChangeSKUHandler(repository Repository) *ChangeSKUHandler {
	return &ChangeSKUHandler{
		repository: repository,
	}
}

func (h *ChangeSKUHandler) Handle(
	ctx context.Context,
	cmd ChangeSKUCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return translateRepositoryError(err)
	}

	if product == nil {
		return ErrProductNotFound
	}

	if product.SKU() != cmd.SKU {
		existing, err := h.repository.FindBySKU(ctx, cmd.SKU)

		if err == nil && existing != nil && existing.ID() != product.ID() {
			return ErrSKUAlreadyExists
		}

		if err != nil && !errors.Is(err, ErrRepositoryNotFound) {
			return err
		}
	}

	if err := product.ChangeSKU(cmd.SKU); err != nil {
		return err
	}

	return h.repository.Update(ctx, product)
}
