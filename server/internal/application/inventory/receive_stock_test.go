package inventory

import (
	"context"
	"errors"
	"testing"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type receiveStockRepository struct {
	inventory *domaininventory.Inventory

	findErr   error
	saveErr   error
	updateErr error

	saveCalls   int
	updateCalls int

	savedInventory   *domaininventory.Inventory
	updatedInventory *domaininventory.Inventory
}

func (r *receiveStockRepository) Save(
	_ context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.saveCalls++
	r.savedInventory = inventory

	return r.saveErr
}

func (r *receiveStockRepository) Update(
	_ context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalls++
	r.updatedInventory = inventory

	return r.updateErr
}

func (r *receiveStockRepository) FindByProductID(
	_ context.Context,
	_ ids.ID,
) (*domaininventory.Inventory, error) {
	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func TestReceiveStockHandler_ExistingInventory(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	initialQuantity, err := quantity.New(10)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReceiveStock(initialQuantity); err != nil {
		t.Fatal(err)
	}

	receivedQuantity, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	repository := &receiveStockRepository{
		inventory: inventory,
	}

	handler := NewReceiveStockHandler(repository)

	err = handler.Handle(context.Background(), ReceiveStockCommand{
		ProductID: productID,
		Quantity:  receivedQuantity,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repository.updateCalls != 1 {
		t.Fatalf("expected 1 update call, got %d", repository.updateCalls)
	}

	if repository.saveCalls != 0 {
		t.Fatalf("expected 0 save calls, got %d", repository.saveCalls)
	}

	if got := inventory.OnHand().Value(); got != 15 {
		t.Fatalf("expected on-hand quantity 15, got %d", got)
	}
}

func TestReceiveStockHandler_MissingInventory(t *testing.T) {
	productID := ids.New()

	receivedQuantity, err := quantity.New(20)
	if err != nil {
		t.Fatal(err)
	}

	repository := &receiveStockRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewReceiveStockHandler(repository)

	err = handler.Handle(context.Background(), ReceiveStockCommand{
		ProductID: productID,
		Quantity:  receivedQuantity,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repository.saveCalls != 1 {
		t.Fatalf("expected 1 save call, got %d", repository.saveCalls)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}

	if repository.savedInventory == nil {
		t.Fatal("expected inventory to be saved")
	}

	if got := repository.savedInventory.ProductID(); got != productID {
		t.Fatalf("expected product ID %v, got %v", productID, got)
	}

	if got := repository.savedInventory.OnHand().Value(); got != 20 {
		t.Fatalf("expected on-hand quantity 20, got %d", got)
	}
}

func TestReceiveStockHandler_InvalidQuantity(t *testing.T) {
	productID := ids.New()

	inventory, err := domaininventory.New(productID)
	if err != nil {
		t.Fatal(err)
	}

	repository := &receiveStockRepository{
		inventory: inventory,
	}

	handler := NewReceiveStockHandler(repository)

	zero, err := quantity.New(0)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReceiveStockCommand{
		ProductID: productID,
		Quantity:  zero,
	})
	if !errors.Is(err, domaininventory.ErrInvalidReceiveQuantity) {
		t.Fatalf(
			"expected ErrInvalidReceiveQuantity, got %v",
			err,
		)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}

	if repository.saveCalls != 0 {
		t.Fatalf("expected 0 save calls, got %d", repository.saveCalls)
	}
}

func TestReceiveStockHandler_RepositoryFailure(t *testing.T) {
	productID := ids.New()

	repositoryError := errors.New("database unavailable")

	repository := &receiveStockRepository{
		findErr: repositoryError,
	}

	handler := NewReceiveStockHandler(repository)

	receivedQuantity, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReceiveStockCommand{
		ProductID: productID,
		Quantity:  receivedQuantity,
	})
	if !errors.Is(err, repositoryError) {
		t.Fatalf("expected repository error, got %v", err)
	}

	if repository.saveCalls != 0 {
		t.Fatalf("expected 0 save calls, got %d", repository.saveCalls)
	}

	if repository.updateCalls != 0 {
		t.Fatalf("expected 0 update calls, got %d", repository.updateCalls)
	}
}

func TestReceiveStockHandler_InvalidProductID(t *testing.T) {
	repository := &receiveStockRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewReceiveStockHandler(repository)

	receivedQuantity, err := quantity.New(5)
	if err != nil {
		t.Fatal(err)
	}

	err = handler.Handle(context.Background(), ReceiveStockCommand{
		ProductID: ids.ID(""),
		Quantity:  receivedQuantity,
	})
	if !errors.Is(err, domaininventory.ErrEmptyProductID) {
		t.Fatalf("expected ErrEmptyProductID, got %v", err)
	}

	if repository.saveCalls != 0 {
		t.Fatalf("expected 0 save calls, got %d", repository.saveCalls)
	}
}
