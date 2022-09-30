package manager

import (
	"errors"
	"fmt"
	"math/big"
	"obra-blanca/event"
	"obra-blanca/event/product"
	"testing"
	"time"
)

func TestManager_CreateEvent(t *testing.T) {
	testManager := createEventManager()

	testEvent := createEvent(t, testManager)

	assertExpectedEventName(t, testEvent)
	assertExpectedEventDate(t, testEvent)
}

func TestEventsManager_AddProductToEvent(t *testing.T) {
	testManager := createEventManager()
	testEvent := createEvent(t, testManager)
	addProductToEvent(&testEvent)

	assertProductWasAddedToEvent(t, testEvent)
}

func assertProductWasAddedToEvent(t *testing.T, testEvent event.DTO) {
	testEventProduct, exists := testEvent.Catalog()[expectedProductName]
	if !exists {
		t.Fatal("product not added to event")
	}
	if n := testEventProduct.Name(); n != expectedProductName {
		t.Fatalf("unexpected product name added to event: %s\n", n)
	}
	if u := testEventProduct.Unit(); u != expectedProductUnit {
		t.Fatalf("unexpected product unit added to event: %s\n", u)
	}
}

func addProductToEvent(testEvent event.Event) {
	var testProduct product.Product = product.New(expectedProductName, expectedProductUnit, product.InventoryDTO{})
	testEvent.AddProduct(testProduct)
}

func createEvent(t *testing.T, testManager eventsManager) event.DTO {
	var err error = testManager.createEvent(expectedEventDate, expectedEventName)
	if err != nil {
		t.Fatalf("error returned while creating event: %s\n", err.Error())
	}
	testEvent, err := testManager.getEvent(expectedEventName)
	return testEvent
}

func assertExpectedEventDate(t *testing.T, testEvent event.DTO) {
	testDate := testEvent.Time()
	if !testDate.Equal(expectedEventDate) {
		t.Fatalf("unexpected date: %s\n", testDate.String())
	}
}

func assertExpectedEventName(t *testing.T, testEvent event.DTO) {
	testName := testEvent.Name()
	if testName != expectedEventName {
		t.Fatalf("unexpected test name: %s\n", testName)
	}
}

func createEventManager() eventsManager {
	var fakePersistence persistence = &stubPersistence{events: map[event.Name]event.DTO{}}
	testManager := getManager(fakePersistence)
	return testManager
}

type stubPersistence struct {
	events map[event.Name]event.DTO
}

func (s *stubPersistence) logNewEvent(name event.Name, t time.Time) error {
	s.events[name] = event.New(t, name)
	return nil
}

func (s *stubPersistence) getLoggedEvent(name event.Name) (event.DTO, error) {
	if _, exists := s.events[name]; !exists {
		return event.DTO{}, errors.New(fmt.Sprintf("event does not exist: %s\n", name))
	}
	return s.events[name], nil
}

type fakeInventory struct {
}

func (f fakeInventory) GetAmount() *big.Float {
	//TODO implement me
	panic("implement me")
}

func (f fakeInventory) SetAside(amount *big.Float) error {
	//TODO implement me
	panic("implement me")
}

const expectedEventName event.Name = "test-name"
const expectedProductName product.Name = "test-product"
const expectedProductUnit product.Unit = "test-unit"

var expectedEventDate time.Time = time.Now()
