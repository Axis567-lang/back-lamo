package persistence

import (
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"io"
	"obra-blanca/event"
	"obra-blanca/event/handler/dto"
	"strings"
)

func New(fileName file.Name) (persistence, error) {
	const persistenceFileName = "eventos.ob"
	if strings.TrimSpace(string(fileName)) == "" {
		fileName = persistenceFileName
	}

	return persistence{
		fileName: fileName,
	}, nil
}

type persistence struct {
	fileName file.Name
}

func (p *persistence) SaveEventDto(event *event.DTO) error {
	var dtos dto.DTOs = make(dto.DTOs)
	err := file.New(p.fileName).Get(&dtos)

	if err != nil && err != io.EOF {
		return fmt.Errorf("could not read existing DTOs: %w", err)
	}

	dtos[event.EventName] = *event
	err = file.New(p.fileName).Save(dtos)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not save DTOs: %w", err)
	}

	return nil
}

func (p *persistence) GetEventDTOs() (dto.DTOs, error) {
	var dtos dto.DTOs
	err := file.New(p.fileName).Get(&dtos)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("could not read existing DTOs: %w", err)
	}
	return dtos, nil
}
