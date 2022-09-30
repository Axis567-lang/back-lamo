package event

import (
	"obra-blanca/distributor"
	"obra-blanca/event/product"
	"time"
)

type Event interface {
	Catalog() Catalog
	Assignments() Assignments
	Name() Name
	Time() time.Time
	AddProduct(product.Product)
	Assign(distributor.Distributor) Assigner
	Status() bool
	Location() Location
	Annotation() Annotation
}

func New(t time.Time, n Name) DTO {
	return DTO{
		EventName:        n,
		EventTime:        t,
		EventCatalog:     make(Catalog),
		EventAssignments: make(Assignments),
	}
}

func (e *DTO) Annotation() Annotation {
	return e.EventAnnotation
}

func (e *DTO) Location() Location {
	return e.EventLocation
}

func (e *DTO) Status() bool {
	return e.EventStatus
}

func (e *DTO) Catalog() Catalog {
	return e.EventCatalog
}

func (e *DTO) Assignments() Assignments {
	return e.EventAssignments
}

func (e *DTO) Assign(d distributor.Distributor) Assigner {
	return Assigner{
		distributor: d,
		event:       e,
	}
}

func (e *DTO) AddProduct(p product.Product) {
	e.EventCatalog[p.Name()] = product.DTO{
		ProductName: p.Name(),
		ProductInventory: product.InventoryDTO{
			InitialAmount: 0,
			Aside:         0,
		},
		ProductUnit: p.Unit(),
	}
}

func (e *DTO) Time() time.Time {
	return e.EventTime
}

func (e *DTO) Name() Name {
	return e.EventName
}

type Catalog map[product.Name]product.DTO

type Name string
type Location string
type Annotation string
