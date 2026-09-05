package inventory

import (
	"context"
	"errors"
	"testing"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type dispatchDirectStockRepository struct {
	inventory        *domaininventory.Inventory
	findErr          error
	updateErr        error
	findCalled       int
	updateCalled     int
	updatedInventory *domaininventory.Inventory
}

func (r *dispatchDirectStockRepository) Save(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	return nil
}

func (r *dispatchDirectStockRepository) Update(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalled++
	r.updatedInventory = inventory

	return r.updateErr
}

func (r *dispatchDirectStockRepository) FindByProductID(
	ctx context.Context,
	productID ids.ID,
) (*domaininventory.Inventory, error) {
	r.findCalled++

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func newDispatchDirectStockInventory(t *testing.T) *domaininventory.Inventory {
	t.Helper()

	inventory, err := domaininventory.New(ids.New())
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(40)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	return inventory
}

func TestDispatchDirectStockHandlerSuccess(t *testing.T) {
	inventory := newDispatchDirectStockInventory(t)

	repository := &dispatchDirectStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: inventory.ProductID(),
			Quantity:  dispatch,
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

	if inventory.OnHand().Value() != 70 {
		t.Fatalf("expected on hand 70, got %d", inventory.OnHand().Value())
	}

	if inventory.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40, got %d", inventory.Reserved().Value())
	}

	available, err := inventory.Available()
	if err != nil {
		t.Fatal(err)
	}

	if available.Value() != 30 {
		t.Fatalf("expected available 30, got %d", available.Value())
	}
}

func TestDispatchDirectStockHandlerRepositoryNotFound(t *testing.T) {
	repository := &dispatchDirectStockRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: ids.New(),
			Quantity:  dispatch,
		},
	)

	if err != ErrInventoryNotFound {
		t.Fatalf("expected %v, got %v", ErrInventoryNotFound, err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchDirectStockHandlerNilInventory(t *testing.T) {
	repository := &dispatchDirectStockRepository{}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: ids.New(),
			Quantity:  dispatch,
		},
	)

	if err != ErrInventoryNotFound {
		t.Fatalf("expected %v, got %v", ErrInventoryNotFound, err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchDirectStockHandlerInsufficientAvailableStock(t *testing.T) {
	inventory := newDispatchDirectStockInventory(t)

	repository := &dispatchDirectStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(70)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: inventory.ProductID(),
			Quantity:  dispatch,
		},
	)

	if err != domaininventory.ErrInsufficientAvailableStock {
		t.Fatalf(
			"expected %v, got %v",
			domaininventory.ErrInsufficientAvailableStock,
			err,
		)
	}

	if inventory.OnHand().Value() != 100 {
		t.Fatalf("expected on hand 100, got %d", inventory.OnHand().Value())
	}

	if inventory.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40, got %d", inventory.Reserved().Value())
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchDirectStockHandlerZeroQuantity(t *testing.T) {
	inventory := newDispatchDirectStockInventory(t)

	repository := &dispatchDirectStockRepository{
		inventory: inventory,
	}

	handler := NewDispatchDirectStockHandler(repository)

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: inventory.ProductID(),
			Quantity:  zero,
		},
	)

	if err != domaininventory.ErrInvalidDirectDispatchQuantity {
		t.Fatalf(
			"expected %v, got %v",
			domaininventory.ErrInvalidDirectDispatchQuantity,
			err,
		)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}

	if inventory.OnHand().Value() != 100 {
		t.Fatalf("expected on hand 100, got %d", inventory.OnHand().Value())
	}

	if inventory.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40, got %d", inventory.Reserved().Value())
	}
}

func TestDispatchDirectStockHandlerFindFailure(t *testing.T) {
	findErr := errors.New("database unavailable")

	repository := &dispatchDirectStockRepository{
		findErr: findErr,
	}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: ids.New(),
			Quantity:  dispatch,
		},
	)

	if !errors.Is(err, findErr) {
		t.Fatalf("expected %v, got %v", findErr, err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d calls", repository.updateCalled)
	}
}

func TestDispatchDirectStockHandlerUpdateFailure(t *testing.T) {
	inventory := newDispatchDirectStockInventory(t)

	updateErr := errors.New("database write failed")

	repository := &dispatchDirectStockRepository{
		inventory: inventory,
		updateErr: updateErr,
	}

	handler := NewDispatchDirectStockHandler(repository)

	dispatch, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(
		context.Background(),
		DispatchDirectStockCommand{
			ProductID: inventory.ProductID(),
			Quantity:  dispatch,
		},
	)

	if !errors.Is(err, updateErr) {
		t.Fatalf("expected %v, got %v", updateErr, err)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update to be called once, got %d", repository.updateCalled)
	}

	if inventory.OnHand().Value() != 70 {
		t.Fatalf("expected on hand 70 after domain mutation, got %d", inventory.OnHand().Value())
	}

	if inventory.Reserved().Value() != 40 {
		t.Fatalf("expected reserved 40 after domain mutation, got %d", inventory.Reserved().Value())
	}
}
