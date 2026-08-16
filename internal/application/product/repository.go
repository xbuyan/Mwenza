package product

import (
	"context"
	"errors"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

var ErrRepositoryNotFound = errors.New("repository entity not found")

type Repository interface {
	Save(ctx context.Context, product *domainproduct.Product) error
	Update(ctx context.Context, product *domainproduct.Product) error
	Delete(ctx context.Context, id ids.ID) error

	FindByID(ctx context.Context, id ids.ID) (*domainproduct.Product, error)
	FindBySKU(ctx context.Context, sku domainproduct.SKU) (*domainproduct.Product, error)
}
