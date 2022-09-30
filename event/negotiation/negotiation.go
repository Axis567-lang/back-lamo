package negotiation

import (
	"obra-blanca/event/product"
)

func New() *DTO {
	return &DTO{NegotiationProducts: make(Products)}
}

type Negotiation interface {
	Products() Products
	AddProduct(product.Product, product.Name)
}

type DTO struct {
	NegotiationProducts Products
}

func (n *DTO) AddProduct(p product.Product, name product.Name) {
	n.NegotiationProducts[name] = product.DTO{
		ProductName:      p.Name(),
		ProductInventory: p.Inventory(),
		ProductUnit:      p.Unit(),
	}
}

func (n *DTO) Products() Products {
	return n.NegotiationProducts
}

type Products map[product.Name]product.DTO
