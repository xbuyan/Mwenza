package inventory

import (
	"context"
	"errors"
	"testing"

	domaininventory "github.com/mwenza/mwenza/internal/domain/inventory"
	inventoryevents "github.com/mwenza/mwenza/internal/domain/inventory/events"
	"github.com/mwenza/mwenza/internal/platform/ids"
	"github.com/mwenza/mwenza/internal/platform/shared/quantity"
)

type stockCountRepository struct {
	inventory        *domaininventory.Inventory
	findErr          error
	updateErr        error
	findCalled       int
	updateCalled     int
	updatedInventory *domaininventory.Inventory
}

func (r *stockCountRepository) Save(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	return nil
}

func (r *stockCountRepository) Update(
	ctx context.Context,
	inventory *domaininventory.Inventory,
) error {
	r.updateCalled++
	r.updatedInventory = inventory
	return r.updateErr
}

func (r *stockCountRepository) FindByProductID(
	ctx context.Context,
	productID ids.ID,
) (*domaininventory.Inventory, error) {
	r.findCalled++

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.inventory, nil
}

func newStockCountInventory(t *testing.T) *domaininventory.Inventory {
	t.Helper()

	inventory, err := domaininventory.New(ids.New())
	if err != nil {
		t.Fatal(err)
	}

	stock, err := quantity.New(100)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReceiveStock(stock); err != nil {
		t.Fatal(err)
	}

	reserved, err := quantity.New(30)
	if err != nil {
		t.Fatal(err)
	}

	if err := inventory.ReserveStock(reserved); err != nil {
		t.Fatal(err)
	}

	// Remove setup events so the test can inspect only the stock count event.
	inventory.Pull()

	return inventory
}

func TestStockCountHandler(t *testing.T) {
	inventory := newStockCountInventory(t)

	counted, err := quantity.New(72)
	if err != nil {
		t.Fatal(err)
	}

	repository := &stockCountRepository{
		inventory: inventory,
	}

	handler := NewStockCountHandler(repository)

	err = handler.Handle(context.Background(), StockCountCommand{
		ProductID: inventory.ProductID(),
		Quantity:  counted,
	})
	if err != nil {
		t.Fatal(err)
	}

	if repository.findCalled != 1 {
		t.Fatalf("expected FindByProductID once, got %d", repository.findCalled)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update once, got %d", repository.updateCalled)
	}

	if repository.updatedInventory != inventory {
		t.Fatal("expected updated inventory to be persisted")
	}

	if inventory.OnHand().Value() != 72 {
		t.Fatalf(
			"expected on hand 72, got %d",
			inventory.OnHand().Value(),
		)
	}

	if inventory.Reserved().Value() != 30 {
		t.Fatalf(
			"expected reserved 30, got %d",
			inventory.Reserved().Value(),
		)
	}

	events := inventory.Pull()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event, ok := events[0].(inventoryevents.StockCounted)
	if !ok {
		t.Fatalf("expected StockCounted event, got %T", events[0])
	}

	if !event.Quantity.Equal(counted) {
		t.Fatalf(
			"expected counted quantity %d, got %d",
			counted.Value(),
			event.Quantity.Value(),
		)
	}
}

func TestStockCountHandlerRepositoryNotFound(t *testing.T) {
	repository := &stockCountRepository{
		findErr: ErrRepositoryNotFound,
	}

	handler := NewStockCountHandler(repository)

	err := handler.Handle(context.Background(), StockCountCommand{
		ProductID: ids.New(),
		Quantity:  mustQuantity(t, 50),
	})

	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected %v, got %v", ErrInventoryNotFound, err)
	}

	if repository.findCalled != 1 {
		t.Fatalf("expected FindByProductID once, got %d", repository.findCalled)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d", repository.updateCalled)
	}
}

func TestStockCountHandlerNilInventory(t *testing.T) {
	repository := &stockCountRepository{}

	handler := NewStockCountHandler(repository)

	err := handler.Handle(context.Background(), StockCountCommand{
		ProductID: ids.New(),
		Quantity:  mustQuantity(t, 50),
	})

	if !errors.Is(err, ErrInventoryNotFound) {
		t.Fatalf("expected %v, got %v", ErrInventoryNotFound, err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d", repository.updateCalled)
	}
}

func TestStockCountHandlerCountBelowReserved(t *testing.T) {
	inventory := newStockCountInventory(t)

	repository := &stockCountRepository{
		inventory: inventory,
	}

	handler := NewStockCountHandler(repository)

	err := handler.Handle(context.Background(), StockCountCommand{
		ProductID: inventory.ProductID(),
		Quantity:  mustQuantity(t, 20),
	})

	if !errors.Is(err, domaininventory.ErrCountBelowReserved) {
		t.Fatalf(
			"expected %v, got %v",
			domaininventory.ErrCountBelowReserved,
			err,
		)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d", repository.updateCalled)
	}

	if inventory.OnHand().Value() != 100 {
		t.Fatalf(
			"expected on hand to remain 100, got %d",
			inventory.OnHand().Value(),
		)
	}

	if inventory.Reserved().Value() != 30 {
		t.Fatalf(
			"expected reserved to remain 30, got %d",
			inventory.Reserved().Value(),
		)
	}

	if inventory.HasEvents() {
		t.Fatal("expected no event for rejected stock count")
	}
}

func TestStockCountHandlerFindFailure(t *testing.T) {
	findErr := errors.New("database unavailable")

	repository := &stockCountRepository{
		findErr: findErr,
	}

	handler := NewStockCountHandler(repository)

	err := handler.Handle(context.Background(), StockCountCommand{
		ProductID: ids.New(),
		Quantity:  mustQuantity(t, 50),
	})

	if !errors.Is(err, findErr) {
		t.Fatalf("expected %v, got %v", findErr, err)
	}

	if repository.updateCalled != 0 {
		t.Fatalf("expected Update not to be called, got %d", repository.updateCalled)
	}
}

func TestStockCountHandlerUpdateFailure(t *testing.T) {
	inventory := newStockCountInventory(t)

	updateErr := errors.New("database write failed")

	repository := &stockCountRepository{
		inventory: inventory,
		updateErr: updateErr,
	}

	handler := NewStockCountHandler(repository)

	err := handler.Handle(context.Background(), StockCountCommand{
		ProductID: inventory.ProductID(),
		Quantity:  mustQuantity(t, 72),
	})

	if !errors.Is(err, updateErr) {
		t.Fatalf("expected %v, got %v", updateErr, err)
	}

	if repository.updateCalled != 1 {
		t.Fatalf("expected Update once, got %d", repository.updateCalled)
	}

	if inventory.OnHand().Value() != 72 {
		t.Fatalf(
			"expected domain mutation to occur before persistence failure, got %d",
			inventory.OnHand().Value(),
		)
	}
}

func mustQuantity(t *testing.T, value int64) quantity.Quantity {
	t.Helper()

	qty, err := quantity.New(value)
	if err != nil {
		t.Fatal(err)
	}

	return qty
}
