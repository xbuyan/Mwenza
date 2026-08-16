package product

import (
	"context"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type ChangeUnitCommand struct {
	ID   ids.ID
	Unit domainproduct.Unit
}

type ChangeUnitHandler struct {
	repository Repository
}

func NewChangeUnitHandler(repository Repository) *ChangeUnitHandler {
	return &ChangeUnitHandler{
		repository: repository,
	}
}

func (h *ChangeUnitHandler) Handle(
	ctx context.Context,
	cmd ChangeUnitCommand,
) error {
	product, err := h.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	if err := product.ChangeUnit(cmd.Unit); err != nil {
		return err
	}

	return h.repository.Update(ctx, product)
}
