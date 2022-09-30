package event

import (
	"math/big"
	"obra-blanca/distributor"
	"obra-blanca/event/negotiation"
	"obra-blanca/event/product"
	"obra-blanca/event/product/inventory"
	"obra-blanca/promoter"
	"testing"
	"time"
)

type fakeNegotiation struct {
	distributor distributor.Distributor
}

func (f fakeNegotiation) Products() negotiation.Products {
	//TODO implement me
	panic("implement me")
}

func (f fakeNegotiation) AddProduct(p product.Product, name product.Name) {
	//TODO implement me
	panic("implement me")
}

func TestDirector_CreateAssignment(t *testing.T) {
	var testDirector Director = CreateDirector(time.Now(), testEventName)

	const testDistributor distributor.Name = "distributor"
	testDirector.CreateNegotiation(testDistributor)
	if _, exists := testDirector.Event().Assignments()[testDistributor]; !exists {
		t.Fatal("assignment not found")
	}
}

func validateTheProductsAmount(t *testing.T, testCatalog Catalog) {
	if testCatalog[testProductName].Inventory().GetAmount().Cmp(big.NewFloat(testAmount)) != 0 {
		retrievedAmount, _ := testCatalog[testProductName].Inventory().GetAmount().Float32()
		t.Fatalf("amounts do not match: expected (%f) -- got (%f)\n", testAmount, retrievedAmount)
	}
}

func checkIfTheProductExists(t *testing.T, testCatalog Catalog) {
	if _, exists := testCatalog[testProductName]; !exists {
		t.Fatalf("expected product was not found in the event")
	}
}

func getTestCatalog(testDirector Director) Catalog {
	var testEvent Event = testDirector.Event()
	var testCatalog = testEvent.Catalog()
	return testCatalog
}

func makeTestInventory() product.Inventory {
	var amt = big.NewFloat(testAmount)
	var testInventory = inventory.New(amt)
	return testInventory
}

type fakePromoter struct {
}

func (f fakePromoter) Name() promoter.Name {
	//TODO implement me
	panic("implement me")
}

type fakeDistributor struct {
}

func (f fakeDistributor) Negotiations() distributor.Negotiations {
	//TODO implement me
	panic("implement me")
}

func (f fakeDistributor) Name() distributor.Name {
	return ""
}

const testProductName product.Name = "test product"
const testAmount = 1000.0
