package inventory

import (
	"fmt"
	"math/big"
	"obra-blanca/event/product"
)

func New(amt *big.Float) product.Inventory {
	return inventory{
		amount: amt,
	}
}

type inventory struct {
	amount *big.Float
}

func (i inventory) SetAside(amount *big.Float) error {
	if i.amount.Cmp(amount) < 1 {
		return FormatInsufficientAmountError(amount, i.amount)
	}
	i.amount.Sub(i.amount, amount)
	return nil
}

func (i inventory) GetAmount() *big.Float {
	return i.amount
}

func FormatInsufficientAmountError(requestedAmount *big.Float, remainingAmount *big.Float) error {
	return fmt.Errorf("cannot set aside %f, there is only %f left", requestedAmount, remainingAmount)
}
