package product

import (
	"context"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

type CreateCommand struct {
	SKU  domainproduct.SKU
	Name string
	Unit domainproduct.Unit
}

func (s *Service) Create(
	ctx context.Context,
	cmd CreateCommand,
) (*domainproduct.Product, error) {
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

	if err := s.repository.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) Get(
	ctx context.Context,
	id ids.ID,
) (*domainproduct.Product, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *Service) GetBySKU(
	ctx context.Context,
	sku domainproduct.SKU,
) (*domainproduct.Product, error) {
	return s.repository.FindBySKU(ctx, sku)
}

func (s *Service) Rename(
	ctx context.Context,
	id ids.ID,
	name string,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := product.Rename(name); err != nil {
		return err
	}

	return s.repository.Update(ctx, product)
}

func (s *Service) ChangeSKU(
	ctx context.Context,
	id ids.ID,
	sku domainproduct.SKU,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := product.ChangeSKU(sku); err != nil {
		return err
	}

	return s.repository.Update(ctx, product)
}

func (s *Service) ChangeUnit(
	ctx context.Context,
	id ids.ID,
	unit domainproduct.Unit,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := product.ChangeUnit(unit); err != nil {
		return err
	}

	return s.repository.Update(ctx, product)
}

func (s *Service) ChangeDescription(
	ctx context.Context,
	id ids.ID,
	description string,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	product.ChangeDescription(description)

	return s.repository.Update(ctx, product)
}

func (s *Service) Activate(
	ctx context.Context,
	id ids.ID,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	product.Activate()

	return s.repository.Update(ctx, product)
}

func (s *Service) Deactivate(
	ctx context.Context,
	id ids.ID,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	product.Deactivate()

	return s.repository.Update(ctx, product)
}

func (s *Service) Discontinue(
	ctx context.Context,
	id ids.ID,
) error {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	product.Discontinue()

	return s.repository.Update(ctx, product)
}
