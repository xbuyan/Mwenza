package inventory

import (
	"context"
	"errors"
	"testing"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type reserveStockRepository struct {
	inventory *domaininventory.Inventory

	findErr   error
	updateErr error

	updateCalls      int
	updatedInventory *domaininventory.Inventory
}

func (r *reserveStockRepository) Save(
	_ context.Context,
	_ *domaininventory.Inventory,
) error {
	return nil
}

func (r *reserveStockRepository) Update(
	_ context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalls++
	r.updatedInventory = inventory

	return r.updateErr
}

func (r *reserveStockRepository) FindByProductID(
	_ context.Context,
	_ ids.ID,
) (*domaininventory.Inventory, error) {
	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func TestReserveStockHandler_ExistingInventory(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
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

	reserve, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	repository := &reserveStockRepository{
		inventory: inventory,
	}

	handler := NewReserveStockHandler(repository)

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: productID,
		Quantity:  reserve,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repository.updateCalls != 1 {
		t.Fatalf("expected 1 update call, got %d", repository.updateCalls)
	}

	if repository.updatedInventory != inventory {
		t.Fatal("expected updated inventory to be the loaded inventory")
	}

	if got := inventory.Reserved().Value(); got != 30 {
		t.Fatalf("expected reserved quantity 30, got %d", got)
	}

	available, err := inventory.Available()
	if err != nil {
		t.Fatal(err)
	}

	if got := available.Value(); got != 70 {
		t.Fatalf("expected available quantity 70, got %d", got)
	}
}

func TestReserveStockHandler_InventoryNotFound(t *testing.T) {
	repository := &reserveStockRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewReserveStockHandler(repository)

	quantityToReserve, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: ids.New(),
		Quantity:  quantityToReserve,
	})
	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}
}

func TestReserveStockHandler_NilInventory(t *testing.T) {
	repository := &reserveStockRepository{}

	handler := NewReserveStockHandler(repository)

	quantityToReserve, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: ids.New(),
		Quantity:  quantityToReserve,
	})
	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected ErrInventoryNotFound, got %v", err)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}
}

func TestReserveStockHandler_InsufficientAvailableStock(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	reserve, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	repository := &reserveStockRepository{
		inventory: inventory,
	}

	handler := NewReserveStockHandler(repository)

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: productID,
		Quantity:  reserve,
	})
	if !errors.Is(err, domaininventory.ErrInsufficientAvailableStock) {
		t.Fatalf(
			"expected ErrInsufficientAvailableStock, got %v",
			err,
		)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}

	if got := inventory.Reserved().Value(); got != 0 {
		t.Fatalf("expected reserved quantity 0, got %d", got)
	}
}

func TestReserveStockHandler_InvalidQuantity(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	repository := &reserveStockRepository{
		inventory: inventory,
	}

	handler := NewReserveStockHandler(repository)

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: productID,
		Quantity:  zero,
	})
	if !errors.Is(err, domaininventory.ErrInvalidReservationQuantity) {
		t.Fatalf(
			"expected ErrInvalidReservationQuantity, got %v",
			err,
		)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}
}

func TestReserveStockHandler_RepositoryFailure(t *testing.T) {
	repositoryError := errors.New("database unavailable")

	repository := &reserveStockRepository{
		findErr: repositoryError,
	}

	handler := NewReserveStockHandler(repository)

	quantityToReserve, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: ids.New(),
		Quantity:  quantityToReserve,
	})
	if !errors.Is(err, repositoryError) {
		t.Fatalf("expected repository error, got %v", err)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}
}

func TestReserveStockHandler_UpdateFailure(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	received, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReceiveStock(received); err != nil {
		t.Fatal(err)
	}

	updateError := errors.New("database write failed")

	repository := &reserveStockRepository{
		inventory: inventory,
		updateErr: updateError,
	}

	handler := NewReserveStockHandler(repository)

	reserve, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReserveStockCommand{
		ProductID: productID,
		Quantity:  reserve,
	})
	if !errors.Is(err, updateError) {
		t.Fatalf("expected update error, got %v", err)
	}

	if repository.updateCalls != 1 {
		t.Fatalf("expected 1 update call, got %d", repository.updateCalls)
	}

	if got := inventory.Reserved().Value(); got != 5 {
		t.Fatalf("expected reserved quantity 5, got %d", got)
	}
}
