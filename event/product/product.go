package product

func New(n Name, u Unit, i InventoryDTO) Product {
	return basicProduct{
		name:      n,
		unit:      u,
		inventory: i,
	}
}

type Product interface {
	Inventory() InventoryDTO
	Name() Name
	Unit() Unit
}

type basicProduct struct {
	name      Name
	unit      Unit
	inventory InventoryDTO
}

func (p basicProduct) Inventory() InventoryDTO {
	return p.inventory
}

func (p basicProduct) Unit() Unit {
	return p.unit
}

func (p basicProduct) Name() Name {
	return p.name
}

type Unit string
type Name string
