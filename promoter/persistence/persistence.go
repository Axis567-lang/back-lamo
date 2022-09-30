package persistence

import (
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"io"
	"obra-blanca/promoter"
)

func New() promoter.Persistence {
	return persistence{fileName: promoterFileName}
}

type persistence struct {
	fileName file.Name
}

func (per persistence) DeletePromoter(name promoter.Name) error {
	promoters, err := per.GetPromoters()
	if err != nil {
		return fmt.Errorf("could not get promoters: %w", err)
	}
	delete(promoters, name)
	fileAccess := file.New(per.fileName)
	err = fileAccess.Save(promoters)
	if err != nil {
		return fmt.Errorf("could not save promoters to persistence: %w", err)
	}
	return nil
}

func (per persistence) GetPromoters() (promoter.Promoters, error) {
	var promoters promoter.Promoters
	err := file.New(per.fileName).Get(&promoters)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("could not get promoters from persistence: %w", err)
	}
	return promoters, nil
}

func (per persistence) AddPromoter(p promoter.Promoter) error {
	var promoters promoter.Promoters
	fileAccess := file.New(per.fileName)
	err := fileAccess.Get(&promoters)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not get promoters from persistence: %w", err)
	}

	if promoters == nil {
		promoters = make(promoter.Promoters)
	}
	promoters[p.Name()] = promoter.DTO{PromoterName: p.Name()}
	err = fileAccess.Save(promoters)
	if err != nil {
		return fmt.Errorf("could not save promoters to persistence: %w", err)
	}
	return nil
}

const promoterFileName file.Name = "promotores.ob"
