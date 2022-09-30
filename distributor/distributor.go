package distributor

import (
	"obra-blanca/event/negotiation"
)

type Distributor interface {
	Name() Name
	Negotiations() Negotiations
}

type DTO struct {
	DistributorName         Name         `json:"name,omitempty"`
	DistributorNegotiations Negotiations `json:"negotiations,omitempty"`
}

func (d DTO) Name() Name {
	return d.DistributorName
}

func (d DTO) Negotiations() Negotiations {
	return d.DistributorNegotiations
}

func New(name Name) DTO {
	return DTO{
		DistributorName:         name,
		DistributorNegotiations: Negotiations{},
	}
}

type Negotiations []negotiation.Negotiation
type Name string
