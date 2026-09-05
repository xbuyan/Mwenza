package sale

import (
	"context"

	domainsale "github.com/mwenza/mwenza/internal/domain/sale"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type Repository interface {
	Save(ctx context.Context, sale *domainsale.Sale) error
	Update(ctx context.Context, sale *domainsale.Sale) error

	FindByID(ctx context.Context, id ids.ID) (*domainsale.Sale, error)
}
