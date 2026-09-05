package inventory_test

import (
	"context"
	"errors"
	"testing"

	appinventory "github.com/mwenza/mwenza/internal/application/inventory"
	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type releaseReservedStockRepository struct {
	inventory        *domaininventory.Inventory
	findErr          error
	updateErr        error
	findCalled       int
	updateCalled     int
	updatedInventory *domaininventory.Inventory
}

func (r *releaseReservedStockRepository) Save(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	return nil
}

func (r *releaseReservedStockRepository) Update(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalled++
	r.updatedInventory = inventory

	return r.updateErr
}

func (r *releaseReservedStockRepository) FindByProductID(
	ctx context.Context,
	productID ids.ID,
) (*domaininventory.Inventory, error) {
	r.findCalled++

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func TestReleaseReservedStockHandler_Success(t *testing.T) {
	productID := ids.New()

	inv, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(40)
	if err != nil {
		t.Fatal(err)
	}

	release, err := quantity.New(15)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	repository := &releaseReservedStockRepository{
		inventory: inv,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: productID,
			Quantity:  release,
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repository.findCalled != 1 {
		t.Fatalf("expected FindByProductID to be called once, got %d", repository.findCalled)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update to be called once, got %d", repository.updateCalled)
	}

	if inv.OnHand().Value() != 100 {
		t.Fatalf("expected on-hand quantity 100, got %d", inv.OnHand().Value())
	}

	if inv.Reserved().Value() != 25 {
		t.Fatalf("expected reserved quantity 25, got %d", inv.Reserved().Value())
	}

	available, err := inv.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 75 {
		t.Fatalf("expected available quantity 75, got %d", available.Value())
	}
}

func TestReleaseReservedStockHandler_RepositoryNotFound(t *testing.T) {
	productID := ids.New()

	repository := &releaseReservedStockRepository{
		findErr: appinventory.ErrRepositoryNotFound,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	qty, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: productID,
			Quantity:  qty,
		},
	)

	if !errors.Is(err, appinventory.ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestReleaseReservedStockHandler_NilInventory(t *testing.T) {
	repository := &releaseReservedStockRepository{
		inventory: nil,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	qty, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: ids.New(),
			Quantity:  qty,
		},
	)

	if !errors.Is(err, appinventory.ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestReleaseReservedStockHandler_InsufficientReservedStock(t *testing.T) {
	productID := ids.New()

	inv, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	release, err := quantity.New(6)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	repository := &releaseReservedStockRepository{
		inventory: inv,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: productID,
			Quantity:  release,
		},
	)

	if !errors.Is(err, domaininventory.ErrInsufficientReservedStock) {
		t.Fatalf(
			"expected ErrInsufficientReservedStock, got %v",
			err,
		)
	}

	if inv.Reserved().Value() != 5 {
		t.Fatalf("expected reserved quantity to remain 5, got %d", inv.Reserved().Value())
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestReleaseReservedStockHandler_ZeroQuantity(t *testing.T) {
	productID := ids.New()

	inv, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	repository := &releaseReservedStockRepository{
		inventory: inv,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: productID,
			Quantity:  zero,
		},
	)

	if !errors.Is(err, domaininventory.ErrInvalidReleaseQuantity) {
		t.Fatalf(
			"expected ErrInvalidReleaseQuantity, got %v",
			err,
		)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestReleaseReservedStockHandler_FindFailure(t *testing.T) {
	repositoryError := errors.New("database unavailable")

	repository := &releaseReservedStockRepository{
		findErr: repositoryError,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	qty, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: ids.New(),
			Quantity:  qty,
		},
	)

	if !errors.Is(err, repositoryError) {
		t.Fatalf("expected repository error, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestReleaseReservedStockHandler_UpdateFailure(t *testing.T) {
	productID := ids.New()

	inv, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(40)
	if err != nil {
		t.Fatal(err)
	}

	release, err := quantity.New(15)
	if err != nil {
		t.Fatal(err)
	}

	if err := inv.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	if err := inv.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	repositoryError := errors.New("database unavailable")

	repository := &releaseReservedStockRepository{
		inventory: inv,
		updateErr: repositoryError,
	}

	handler := appinventory.NewReleaseReservedStockHandler(repository)

	err = handler.Handle(
		context.Background(),
		appinventory.ReleaseReservedStockCommand{
			ProductID: productID,
			Quantity:  release,
		},
	)

	if !errors.Is(err, repositoryError) {
		t.Fatalf("expected repository error, got %v", err)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update to be called once, got %d", repository.updateCalled)
	}

	if inv.Reserved().Value() != 25 {
		t.Fatalf("expected domain mutation to leave reserved at 25, got %d", inv.Reserved().Value())
	}
}
