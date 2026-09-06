package inventory

import (
	"context"
	"errors"
	"testing"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type dispatchReservedStockRepository struct {
	inventory        *domaininventory.Inventory
	findErr          error
	updateErr        error
	findCalled       int
	updateCalled     int
	updatedInventory *domaininventory.Inventory
}

func (r *dispatchReservedStockRepository) Save(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	return nil
}

func (r *dispatchReservedStockRepository) Update(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalled++
	r.updatedInventory = inventory
	return r.updateErr
}

func (r *dispatchReservedStockRepository) FindByProductID(
	ctx context.Context,
	productID ids.ID,
) (*domaininventory.Inventory, error) {
	r.findCalled++

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func newDispatchReservedStockInventory(t *testing.T) (*domaininventory.Inventory, ids.ID) {
	t.Helper()

	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatalf("failed to create inventory: %v", err)
	}

	received, err := quantity.New(100)
	if err != nil {
		t.Fatalf("failed to create received quantity: %v", err)
	}

	if err := inventory.ReceiveStock(received); err != nil {
		t.Fatalf("failed to receive stock: %v", err)
	}

	reserved, err := quantity.New(40)
	if err != nil {
		t.Fatalf("failed to create reserved quantity: %v", err)
	}

	if err := inventory.ReserveStock(reserved); err != nil {
		t.Fatalf("failed to reserve stock: %v", err)
	}

	return inventory, productID
}

func TestDispatchReservedStockHandler_Handle_Success(t *testing.T) {
	inventory, productID := newDispatchReservedStockInventory(t)

	repository := &dispatchReservedStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: productID,
			Quantity:  qty,
		},
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if repository.findCalled != 1 {
		t.Fatalf("expected FindByProductID to be called once, got %d", repository.findCalled)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update to be called once, got %d", repository.updateCalled)
	}

	if got := repository.updatedInventory.OnHand().Value(); got != 85 {
		t.Fatalf("expected on-hand quantity 85, got %d", got)
	}

	if got := repository.updatedInventory.Reserved().Value(); got != 25 {
		t.Fatalf("expected reserved quantity 25, got %d", got)
	}

	available, err := repository.updatedInventory.Available()
	if err != nil {
		t.Fatalf("failed to calculate available quantity: %v", err)
	}

	if got := available.Value(); got != 60 {
		t.Fatalf("expected available quantity 60, got %d", got)
	}
}

func TestDispatchReservedStockHandler_Handle_RepositoryNotFound(t *testing.T) {
	repository := &dispatchReservedStockRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: ids.New(),
			Quantity:  qty,
		},
	)
	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchReservedStockHandler_Handle_NilInventory(t *testing.T) {
	repository := &dispatchReservedStockRepository{}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: ids.New(),
			Quantity:  qty,
		},
	)
	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchReservedStockHandler_Handle_InsufficientReservedStock(t *testing.T) {
	inventory, productID := newDispatchReservedStockInventory(t)

	repository := &dispatchReservedStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(50)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: productID,
			Quantity:  qty,
		},
	)
	if !errors.Is(err, domaininventory.ErrInsufficientReservedStock) {
		t.Fatalf(
			"expected ErrInsufficientReservedStock, got %v",
			err,
		)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}

	if got := inventory.OnHand().Value(); got != 100 {
		t.Fatalf("expected on-hand quantity to remain 100, got %d", got)
	}

	if got := inventory.Reserved().Value(); got != 40 {
		t.Fatalf("expected reserved quantity to remain 40, got %d", got)
	}
}

func TestDispatchReservedStockHandler_Handle_ZeroQuantity(t *testing.T) {
	inventory, productID := newDispatchReservedStockInventory(t)

	repository := &dispatchReservedStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(0)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: productID,
			Quantity:  qty,
		},
	)
	if !errors.Is(err, domaininventory.ErrInvalidDispatchQuantity) {
		t.Fatalf(
			"expected ErrInvalidDispatchQuantity, got %v",
			err,
		)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}

	if got := inventory.OnHand().Value(); got != 100 {
		t.Fatalf("expected on-hand quantity to remain 100, got %d", got)
	}

	if got := inventory.Reserved().Value(); got != 40 {
		t.Fatalf("expected reserved quantity to remain 40, got %d", got)
	}
}

func TestDispatchReservedStockHandler_Handle_FindFailure(t *testing.T) {
	findErr := errors.New("database unavailable")

	repository := &dispatchReservedStockRepository{
		findErr: findErr,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: ids.New(),
			Quantity:  qty,
		},
	)
	if !errors.Is(err, findErr) {
		t.Fatalf("expected repository error, got %v", err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchReservedStockHandler_Handle_UpdateFailure(t *testing.T) {
	inventory, productID := newDispatchReservedStockInventory(t)

	updateErr := errors.New("database update failed")

	repository := &dispatchReservedStockRepository{
		inventory: inventory,
		updateErr: updateErr,
	}

	handler := NewDispatchReservedStockHandler(repository)

	qty, err := quantity.New(15)
	if err != nil {
		t.Fatalf("failed to create quantity: %v", err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchReservedStockCommand{
			ProductID: productID,
			Quantity:  qty,
		},
	)
	if !errors.Is(err, updateErr) {
		t.Fatalf("expected update error, got %v", err)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update to be called once, got %d", repository.updateCalled)
	}

	if got := inventory.OnHand().Value(); got != 85 {
		t.Fatalf("expected on-hand quantity 85, got %d", got)
	}

	if got := inventory.Reserved().Value(); got != 25 {
		t.Fatalf("expected reserved quantity 25, got %d", got)
	}
}
