package handler

import (
	"obra-blanca/event"
	"obra-blanca/event/handler/dto"
)

type Persistence interface {
	SaveEventDto(e *event.DTO) error
	GetEventDTOs() (dto.DTOs, error)
}
