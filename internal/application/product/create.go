package product

import (
	"context"
	"errors"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type CreateCommand struct {
	SKU         domainproduct.SKU
	Name        string
	Description string
	Unit        domainproduct.Unit
}

type CreateHandler struct {
	repository Repository
}

func NewCreateHandler(repository Repository) *CreateHandler {
	return &CreateHandler{
		repository: repository,
	}
}

func (h *CreateHandler) Handle(
	ctx context.Context,
	cmd CreateCommand,
) (*domainproduct.Product, error) {
	existing, err := h.repository.FindBySKU(ctx, cmd.SKU)
	if err == nil && existing != nil {
		return nil, ErrSKUAlreadyExists
	}

	if err != nil && !errors.Is(err, ErrRepositoryNotFound) {
		return nil, err
	}

	id := ids.New()

	product, err := domainproduct.New(
		id,
		cmd.SKU,
		cmd.Name,
		cmd.Unit,
	)
	if err != nil {
		return nil, err
	}

	if cmd.Description != "" {
		product.ChangeDescription(cmd.Description)
	}

	if err := h.repository.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
