package product

type DTO struct {
	ProductName      Name         `json:"name"`
	ProductInventory InventoryDTO `json:"inventory"`
	ProductUnit      Unit         `json:"unit"`
}

func (d DTO) Inventory() InventoryDTO {
	return d.ProductInventory
}

func (d DTO) Name() Name {
	return d.ProductName
}

func (d DTO) Unit() Unit {
	return d.ProductUnit
}
