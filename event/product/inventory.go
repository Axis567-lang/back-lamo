package product

import "math/big"

type Inventory interface {
	GetAmount() *big.Float
	SetAside(amount *big.Float) error
}

type InventoryDTO struct {
	InitialAmount float64
	Aside         float64
}

func (i InventoryDTO) GetAmount() *big.Float {
	return big.NewFloat(i.InitialAmount)
}

func (i InventoryDTO) SetAside(amount *big.Float) (err error) {
	i.Aside, _ = amount.Add(amount, big.NewFloat(i.Aside)).Float64()
	return nil
}
