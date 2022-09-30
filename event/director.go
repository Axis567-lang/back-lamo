package event

import (
	"obra-blanca/distributor"
	"obra-blanca/event/negotiation"
	"time"
)

type Director interface {
	Event() Event
	CreateNegotiation(distributor.Name)
}

type director struct {
	event *DTO
}

func (d *director) CreateNegotiation(dis distributor.Name) {
	if d.event.EventAssignments == nil {
		d.event.EventAssignments = make(Assignments)
	}
	d.event.EventAssignments[dis] = negotiation.DTO{
		NegotiationProducts: make(negotiation.Products),
	}
}

func (d *director) Event() Event {
	return d.event
}

func CreateDirector(eventTime time.Time, eventName Name) Director {
	e := New(eventTime, eventName)
	return &director{event: &e}
}
