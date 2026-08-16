package product

import (
	"context"
	"errors"
	"testing"

	domainproduct "github.com/mwenza/mwenza/internal/domain/product"
	"github.com/mwenza/mwenza/internal/platform/ids"
)

var errProductNotFound = errors.New("product not found")

type fakeRepository struct {
	products map[ids.ID]*domainproduct.Product

	saveCalls   int
	updateCalls int
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
	r.products[product.ID()] = product
	return nil
}

func (r *fakeRepository) Update(
	_ context.Context,
	product *domainproduct.Product,
) error {
	r.updateCalls++
	r.products[product.ID()] = product
	return nil
}

func (r *fakeRepository) Delete(
	_ context.Context,
	id ids.ID,
) error {
	delete(r.products, id)
	return nil
}

func (r *fakeRepository) FindByID(
	_ context.Context,
	id ids.ID,
) (*domainproduct.Product, error) {
	product, ok := r.products[id]
	if !ok {
		return nil, errProductNotFound
	}

	return product, nil
}

func (r *fakeRepository) FindBySKU(
	_ context.Context,
	sku domainproduct.SKU,
) (*domainproduct.Product, error) {
	for _, product := range r.products {
		if product.SKU() == sku {
			return product, nil
		}
	}

	return nil, errProductNotFound
}

func TestCreateProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Bamburi Cement 50kg",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
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

	if repository.saveCalls != 1 {
		t.Fatalf("expected 1 save call, got %d", repository.saveCalls)
	}
}

func TestCreateProductRejectsInvalidProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	_, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "",
			Name: "Bamburi Cement 50kg",
			Unit: domainproduct.UnitBag,
		},
	)

	if err != domainproduct.ErrEmptySKU {
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

func TestGetProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	found, err := service.Get(context.Background(), product.ID())
	if err != nil {
		t.Fatal(err)
	}

	if !found.Equals(product) {
		t.Fatal("expected retrieved product to equal created product")
	}
}

func TestGetProductBySKU(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	found, err := service.GetBySKU(
		context.Background(),
		"CEM-001",
	)
	if err != nil {
		t.Fatal(err)
	}

	if !found.Equals(product) {
		t.Fatal("expected retrieved product to equal created product")
	}
}

func TestRenameProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Old Name",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Rename(
		context.Background(),
		product.ID(),
		"New Name",
	)
	if err != nil {
		t.Fatal(err)
	}

	if product.Name() != "New Name" {
		t.Fatalf("expected New Name, got %s", product.Name())
	}

	if repository.updateCalls != 1 {
		t.Fatalf("expected 1 update call, got %d", repository.updateCalls)
	}
}

func TestChangeSKU(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.ChangeSKU(
		context.Background(),
		product.ID(),
		"CEM-002",
	)
	if err != nil {
		t.Fatal(err)
	}

	if product.SKU() != "CEM-002" {
		t.Fatalf("expected CEM-002, got %s", product.SKU())
	}
}

func TestChangeUnit(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.ChangeUnit(
		context.Background(),
		product.ID(),
		domainproduct.UnitPiece,
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

func TestChangeDescription(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.ChangeDescription(
		context.Background(),
		product.ID(),
		"Premium Portland Cement",
	)
	if err != nil {
		t.Fatal(err)
	}

	if product.Description() != "Premium Portland Cement" {
		t.Fatalf(
			"unexpected description: %s",
			product.Description(),
		)
	}
}

func TestDeactivateProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Deactivate(
		context.Background(),
		product.ID(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusInactive {
		t.Fatalf(
			"expected %s, got %s",
			domainproduct.StatusInactive,
			product.Status(),
		)
	}
}

func TestActivateProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := service.Deactivate(
		context.Background(),
		product.ID(),
	); err != nil {
		t.Fatal(err)
	}

	if err := service.Activate(
		context.Background(),
		product.ID(),
	); err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusActive {
		t.Fatalf(
			"expected %s, got %s",
			domainproduct.StatusActive,
			product.Status(),
		)
	}
}

func TestDiscontinueProduct(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Discontinue(
		context.Background(),
		product.ID(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusDiscontinued {
		t.Fatalf(
			"expected %s, got %s",
			domainproduct.StatusDiscontinued,
			product.Status(),
		)
	}
}

func TestDiscontinuedProductCannotBeReactivated(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		CreateCommand{
			SKU:  "CEM-001",
			Name: "Cement",
			Unit: domainproduct.UnitBag,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := service.Discontinue(
		context.Background(),
		product.ID(),
	); err != nil {
		t.Fatal(err)
	}

	if err := service.Activate(
		context.Background(),
		product.ID(),
	); err != nil {
		t.Fatal(err)
	}

	if product.Status() != domainproduct.StatusDiscontinued {
		t.Fatalf(
			"expected discontinued status, got %s",
			product.Status(),
		)
	}
}

func TestRepositoryErrorsArePropagated(t *testing.T) {
	repository := newFakeRepository()
	service := NewService(repository)

	_, err := service.Get(
		context.Background(),
		ids.New(),
	)

	if err != errProductNotFound {
		t.Fatalf(
			"expected %v, got %v",
			errProductNotFound,
			err,
		)
	}
}
