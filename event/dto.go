package event

import (
	"time"
)

type DTO struct {
	EventCatalog     Catalog     `json:"catalog,omitempty"`
	EventAssignments Assignments `json:"assignments,omitempty"`
	EventName        Name        `json:"name"`
	EventTime        time.Time   `json:"time"`
	EventStatus      bool        `json:"status"`
	EventLocation    Location    `json:"location,omitempty"`
	EventAnnotation  Annotation  `json:"annotation,omitempty"`
}
