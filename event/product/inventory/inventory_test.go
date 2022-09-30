package inventory

import (
	"fmt"
	"math/big"
	"obra-blanca/event/product"
	"testing"
)

func TestInventory_GetAmount(t *testing.T) {
	testAmount, testInventory := createTestInventory()
	var amt *big.Float = testInventory.GetAmount()
	if amt.Cmp(testAmount) != 0 {
		t.Fatalf("incorrect product amount")
	}
}

func TestInventory_GetUnit(t *testing.T) {
	testAmount, _ := createTestInventory()

	var _ product.Inventory = New(testAmount)
}

func TestInventory_SetAside(t *testing.T) {
	var testAmount = big.NewFloat(100.0)
	var testInventory product.Inventory = New(testAmount)
	var err error = testInventory.SetAside(big.NewFloat(10))
	if err != nil {
		t.Fatalf(err.Error())
	}
	var amt *big.Float = testInventory.GetAmount()
	if amt.Cmp(big.NewFloat(90)) != 0 {
		value, _ := amt.Float32()
		t.Fatalf("incorrect product amount: %f\n", value)
	}
}

func TestInventory_SetAside_NotEnoughError(t *testing.T) {
	var testAmount = big.NewFloat(100.0)
	var testInventory product.Inventory = New(testAmount)
	err := testInventory.SetAside(big.NewFloat(101))
	if err == nil {
		t.Fatal("did not detect invalid amount for setting aside")
	}
	if err.Error() != FormatInsufficientAmountError(big.NewFloat(101), testAmount).Error() {
		t.Fatalf("invalid error returned: %s\n", err.Error())
	}
	fmt.Printf("✔️ correct error returned: %s\n", err)
}

func createTestInventory() (*big.Float, product.Inventory) {
	var testAmount = big.NewFloat(100.0)
	var testInventory product.Inventory = New(testAmount)
	return testAmount, testInventory
}

var _ product.Inventory = New(big.NewFloat(0))

const fakeName product.Name = "fake name"
