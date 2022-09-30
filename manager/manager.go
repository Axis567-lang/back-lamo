package manager

import (
	"obra-blanca/event"
	"time"
)

type eventsManager struct {
	persistence persistence
}

type persistence interface {
	logNewEvent(event.Name, time.Time) error
	getLoggedEvent(event.Name) (event.DTO, error)
}

func (m *eventsManager) createEvent(eventDate time.Time, eventName event.Name) error {
	return m.persistence.logNewEvent(eventName, eventDate)
}

func (m *eventsManager) getEvent(name event.Name) (event.DTO, error) {
	return m.persistence.getLoggedEvent(name)
}

func getManager(p persistence) eventsManager {
	return eventsManager{
		persistence: p,
	}
}
