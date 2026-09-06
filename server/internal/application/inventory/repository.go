package inventory

import (
	"context"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type Repository interface {
	Save(ctx context.Context, inventory *domaininventory.Inventory) error
	Update(ctx context.Context, inventory *domaininventory.Inventory) error

	FindByProductID(ctx context.Context, productID ids.ID) (*domaininventory.Inventory, error)
}
