package product

import (
	"math/big"
	"testing"
)

func TestBasicProduct_Name(t *testing.T) {
	var testProduct Product = New(expectedProductName, "", InventoryDTO{})
	if testProduct.Name() != expectedProductName {
		t.Fatalf("names do not match: expected (%s) -- obtained (%s)\n", expectedProductName, testProduct.Name())
	}
}

func TestBasicProduct_Unit(t *testing.T) {
	var testProduct Product = New("", expectedProductUnit, InventoryDTO{})
	if testProduct.Unit() != expectedProductUnit {
		t.Fatalf("units do not match: expected (%s) -- obtained (%s)\n", expectedProductUnit, testProduct.Unit())
	}
}

var _ Product = basicProduct{}

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

const expectedProductName Name = "test name"
const expectedProductUnit Unit = "test unit"
