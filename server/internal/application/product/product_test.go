package product

import (
	"context"
	"errors"
	"testing"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

type fakeRepository struct {
	products map[ids.ID]*domainproduct.Product

	saveCalls   int
	updateCalls int
	deleteCalls int

	findByIDErr  error
	findBySKUErr error
	saveErr      error
	updateErr    error
	deleteErr    error
}

func newFakeRepository() *fakeRepository {
	return &fakeRepository{
		products: make(map[ids.ID]*domainproduct.Product),
	}
}

func (r *fakeRepository) Save(
	_ context.Context,
	product *domainproduct.Product,
) error {
	r.saveCalls++

	if r.saveErr != nil {
		return r.saveErr
	}

	r.products[product.ID()] = product
	return nil
}

func (r *fakeRepository) Update(
	_ context.Context,
	product *domainproduct.Product,
) error {
	r.updateCalls++

	if r.updateErr != nil {
		return r.updateErr
	}

	r.products[product.ID()] = product
	return nil
}

func (r *fakeRepository) Delete(
	_ context.Context,
	id ids.ID,
) error {
	r.deleteCalls++

	if r.deleteErr != nil {
		return r.deleteErr
	}

	if _, ok := r.products[id]; !ok {
		return ErrRepositoryNotFound
	}

	delete(r.products, id)
	return nil
}

func (r *fakeRepository) FindByID(
	_ context.Context,
	id ids.ID,
) (*domainproduct.Product, error) {
	if r.findByIDErr != nil {
		return nil, r.findByIDErr
	}

	product, ok := r.products[id]
	if !ok {
		return nil, ErrRepositoryNotFound
	}

	return product, nil
}

func (r *fakeRepository) FindBySKU(
	_ context.Context,
	sku domainproduct.SKU,
) (*domainproduct.Product, error) {
	if r.findBySKUErr != nil {
		return nil, r.findBySKUErr
	}

	for _, product := range r.products {
		if product.SKU() == sku {
			return product, nil
		}
	}

	return nil, ErrRepositoryNotFound
}

func newTestProduct(t *testing.T) *domainproduct.Product {
	t.Helper()

	product, err := domainproduct.New(
		ids.New(),
		"CEM-001",
		"Cement",
		domainproduct.UnitBag,
	)
	if err != nil {
		t.Fatal(err)
	}

	return product
}

func TestCreateHandler(t *testing.T) {
	repository := newFakeRepository()
	handler := NewCreateHandler(repository)

	product, err := handler.Handle(
		context.Background(),
		CreateCommand{
			SKU:         "CEM-001",
			Name:        "Bamburi Cement 50kg",
			Description: "Premium Portland Cement",
			Unit:        domainproduct.UnitBag,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product == nil {
		t.Fatal("expected product")
	}

	if product.ID().IsZero() {
		t.Fatal("expected generated product ID")
	}

	if product.SKU() != "CEM-001" {
		t.Fatalf("expected CEM-001, got %s", product.SKU())
	}

	if product.Name() != "Bamburi Cement 50kg" {
		t.Fatalf(
			"expected Bamburi Cement 50kg, got %s",
			product.Name(),
		)
	}

	if product.Description() != "Premium Portland Cement" {
		t.Fatalf(
			"expected description to be set, got %s",
			product.Description(),
		)
	}

	if repository.saveCalls != 1 {
		t.Fatalf("expected 1 save call, got %d", repository.saveCalls)
	}
}

func TestCreateHandlerRejectsDuplicateSKU(t *testing.T) {
	repository := newFakeRepository()

	existing := newTestProduct(t)

	if err := repository.Save(context.Background(), existing); err != nil {
		t.Fatal(err)
	}

	handler := NewCreateHandler(repository)

	_, err := handler.Handle(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Another Cement",
			Unit: domainproduct.UnitBag,
		},
	)

	if !errors.Is(err, ErrSKUAlreadyExists) {
		t.Fatalf(
			"expected %v, got %v",
			ErrSKUAlreadyExists,
			err,
		)
	}
}

func TestCreateHandlerRejectsInvalidProduct(t *testing.T) {
	repository := newFakeRepository()
	handler := NewCreateHandler(repository)

	_, err := handler.Handle(
		context.Background(),
		CreateCommand{
			SKU:  "",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)

	if !errors.Is(err, domainproduct.ErrEmptySKU) {
		t.Fatalf(
			"expected %v, got %v",
			domainproduct.ErrEmptySKU,
			err,
		)
	}

	if repository.saveCalls != 0 {
		t.Fatal("invalid product must not be saved")
	}
}

func TestRenameHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewRenameHandler(repository)

	err := handler.Handle(
		context.Background(),
		RenameCommand{
			ID:   product.ID(),
			Name: "Premium Cement",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Name() != "Premium Cement" {
		t.Fatalf(
			"expected Premium Cement, got %s",
			product.Name(),
		)
	}

	if repository.updateCalls != 1 {
		t.Fatalf(
			"expected 1 update call, got %d",
			repository.updateCalls,
		)
	}
}

func TestChangeSKUHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewChangeSKUHandler(repository)

	err := handler.Handle(
		context.Background(),
		ChangeSKUCommand{
			ID:  product.ID(),
			SKU: "CEM-002",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.SKU() != "CEM-002" {
		t.Fatalf(
			"expected CEM-002, got %s",
			product.SKU(),
		)
	}
}

func TestChangeSKUHandlerRejectsDuplicateSKU(t *testing.T) {
	repository := newFakeRepository()

	first := newTestProduct(t)

	second, err := domainproduct.New(
		ids.New(),
		"CEM-002",
		"Second Cement",
		domainproduct.UnitBag,
	)
	if err != nil {
		t.Fatal(err)
	}

	repository.products[first.ID()] = first
	repository.products[second.ID()] = second

	handler := NewChangeSKUHandler(repository)

	err = handler.Handle(
		context.Background(),
		ChangeSKUCommand{
			ID:  first.ID(),
			SKU: "CEM-002",
		},
	)

	if !errors.Is(err, ErrSKUAlreadyExists) {
		t.Fatalf(
			"expected %v, got %v",
			ErrSKUAlreadyExists,
			err,
		)
	}

	if first.SKU() != "CEM-001" {
		t.Fatalf(
			"SKU should remain unchanged, got %s",
			first.SKU(),
		)
	}
}

func TestChangeUnitHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewChangeUnitHandler(repository)

	err := handler.Handle(
		context.Background(),
		ChangeUnitCommand{
			ID:   product.ID(),
			Unit: domainproduct.UnitPiece,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Unit() != domainproduct.UnitPiece {
		t.Fatalf(
			"expected %s, got %s",
			domainproduct.UnitPiece,
			product.Unit(),
		)
	}
}

func TestChangeDescriptionHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewChangeDescriptionHandler(repository)

	err := handler.Handle(
		context.Background(),
		ChangeDescriptionCommand{
			ID:          product.ID(),
			Description: "Premium Portland Cement",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Description() != "Premium Portland Cement" {
		t.Fatalf(
			"expected description to change, got %s",
			product.Description(),
		)
	}
}

func TestActivateHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	product.Deactivate()

	repository.products[product.ID()] = product

	handler := NewActivateHandler(repository)

	err := handler.Handle(
		context.Background(),
		ActivateCommand{
			ID: product.ID(),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusActive {
		t.Fatalf(
			"expected active, got %s",
			product.Status(),
		)
	}
}

func TestDeactivateHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewDeactivateHandler(repository)

	err := handler.Handle(
		context.Background(),
		DeactivateCommand{
			ID: product.ID(),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusInactive {
		t.Fatalf(
			"expected inactive, got %s",
			product.Status(),
		)
	}
}

func TestDiscontinueHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewDiscontinueHandler(repository)

	err := handler.Handle(
		context.Background(),
		DiscontinueCommand{
			ID: product.ID(),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusDiscontinued {
		t.Fatalf(
			"expected discontinued, got %s",
			product.Status(),
		)
	}
}

func TestDeleteHandler(t *testing.T) {
	repository := newFakeRepository()
	product := newTestProduct(t)

	repository.products[product.ID()] = product

	handler := NewDeleteHandler(repository)

	err := handler.Handle(
		context.Background(),
		DeleteCommand{
			ID: product.ID(),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if repository.deleteCalls != 1 {
		t.Fatalf(
			"expected 1 delete call, got %d",
			repository.deleteCalls,
		)
	}

	if _, ok := repository.products[product.ID()]; ok {
		t.Fatal("expected product to be deleted")
	}
}

func TestHandlersReturnProductNotFound(t *testing.T) {
	repository := newFakeRepository()
	missingID := ids.New()

	tests := []struct {
		name   string
		handle func() error
	}{
		{
			name: "rename",
			handle: func() error {
				return NewRenameHandler(repository).Handle(
					context.Background(),
					RenameCommand{
						ID:   missingID,
						Name: "New Name",
					},
				)
			},
		},
		{
			name: "change sku",
			handle: func() error {
				return NewChangeSKUHandler(repository).Handle(
					context.Background(),
					ChangeSKUCommand{
						ID:  missingID,
						SKU: "CEM-002",
					},
				)
			},
		},
		{
			name: "change unit",
			handle: func() error {
				return NewChangeUnitHandler(repository).Handle(
					context.Background(),
					ChangeUnitCommand{
						ID:   missingID,
						Unit: domainproduct.UnitPiece,
					},
				)
			},
		},
		{
			name: "change description",
			handle: func() error {
				return NewChangeDescriptionHandler(repository).Handle(
					context.Background(),
					ChangeDescriptionCommand{
						ID:          missingID,
						Description: "Description",
					},
				)
			},
		},
		{
			name: "activate",
			handle: func() error {
				return NewActivateHandler(repository).Handle(
					context.Background(),
					ActivateCommand{
						ID: missingID,
					},
				)
			},
		},
		{
			name: "deactivate",
			handle: func() error {
				return NewDeactivateHandler(repository).Handle(
					context.Background(),
					DeactivateCommand{
						ID: missingID,
					},
				)
			},
		},
		{
			name: "discontinue",
			handle: func() error {
				return NewDiscontinueHandler(repository).Handle(
					context.Background(),
					DiscontinueCommand{
						ID: missingID,
					},
				)
			},
		},
		{
			name: "delete",
			handle: func() error {
				return NewDeleteHandler(repository).Handle(
					context.Background(),
					DeleteCommand{
						ID: missingID,
					},
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.handle()

			if !errors.Is(err, ErrProductNotFound) {
				t.Fatalf(
					"expected %v, got %v",
					ErrProductNotFound,
					err,
				)
			}
		})
	}
}

func TestHandlersPropagateRepositoryErrors(t *testing.T) {
	repositoryError := errors.New("database connection failed")

	t.Run("rename", func(t *testing.T) {
		repository := newFakeRepository()
		repository.findByIDErr = repositoryError

		err := NewRenameHandler(repository).Handle(
			context.Background(),
			RenameCommand{
				ID:   ids.New(),
				Name: "New Name",
			},
		)

		if !errors.Is(err, repositoryError) {
			t.Fatalf(
				"expected repository error %v, got %v",
				repositoryError,
				err,
			)
		}
	})

	t.Run("create", func(t *testing.T) {
		repository := newFakeRepository()
		repository.findBySKUErr = repositoryError

		_, err := NewCreateHandler(repository).Handle(
			context.Background(),
			CreateCommand{
				SKU:  "CEM-001",
				Name: "Cement",
				Unit: domainproduct.UnitBag,
			},
		)

		if !errors.Is(err, repositoryError) {
			t.Fatalf(
				"expected repository error %v, got %v",
				repositoryError,
				err,
			)
		}
	})

	t.Run("change sku", func(t *testing.T) {
		repository := newFakeRepository()
		product := newTestProduct(t)

		repository.products[product.ID()] = product
		repository.findBySKUErr = repositoryError

		err := NewChangeSKUHandler(repository).Handle(
			context.Background(),
			ChangeSKUCommand{
				ID:  product.ID(),
				SKU: "CEM-002",
			},
		)

		if !errors.Is(err, repositoryError) {
			t.Fatalf(
				"expected repository error %v, got %v",
				repositoryError,
				err,
			)
		}
	})

	t.Run("update", func(t *testing.T) {
		repository := newFakeRepository()
		product := newTestProduct(t)

		repository.products[product.ID()] = product
		repository.updateErr = repositoryError

		err := NewRenameHandler(repository).Handle(
			context.Background(),
			RenameCommand{
				ID:   product.ID(),
				Name: "New Name",
			},
		)

		if !errors.Is(err, repositoryError) {
			t.Fatalf(
				"expected repository error %v, got %v",
				repositoryError,
				err,
			)
		}
	})

	t.Run("delete", func(t *testing.T) {
		repository := newFakeRepository()
		product := newTestProduct(t)

		repository.products[product.ID()] = product
		repository.deleteErr = repositoryError

		err := NewDeleteHandler(repository).Handle(
			context.Background(),
			DeleteCommand{
				ID: product.ID(),
			},
		)

		if !errors.Is(err, repositoryError) {
			t.Fatalf(
				"expected repository error %v, got %v",
				repositoryError,
				err,
			)
		}
	})
}
