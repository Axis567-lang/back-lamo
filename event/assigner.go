package event

import (
	"obra-blanca/distributor"
	"obra-blanca/event/negotiation"
)

type Assignments map[distributor.Name]negotiation.DTO

type Assigner struct {
	distributor distributor.Distributor
	event       *DTO
}

func (a Assigner) To(p DTO) {
	p.EventAssignments[a.distributor.Name()] = negotiation.DTO{}
}
