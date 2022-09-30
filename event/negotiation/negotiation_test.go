package negotiation

import (
	"math/big"
	"obra-blanca/event/product"
	"testing"
)

func TestNew(t *testing.T) {
	var _ Negotiation = New()
}

func TestNegotiation(t *testing.T) {
	var testNegotiation Negotiation = New()
	testNegotiation.AddProduct(fakeProduct{}, fakeProductName)
	testProducts := testNegotiation.Products()
	if _, exists := testProducts[fakeProductName]; !exists {
		t.Fatal("did not find added product " + fakeProductName)
	}
}

type fakeProduct struct {
}

func (p fakeProduct) Inventory() product.InventoryDTO {
	return product.InventoryDTO{}
}

func (p fakeProduct) Unit() product.Unit {
	return ""
}

func (p fakeProduct) Name() product.Name {
	return fakeProductName
}

func (p fakeProduct) AvailableAmount() *big.Float {
	//TODO implement me
	panic("implement me")
}

var _ product.Product = fakeProduct{}

const fakeProductName = "fakeProductName"
