package persistence

import (
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"io"
	"obra-blanca/distributor"
)

type persistence struct {
	fileName file.Name
}

func (p persistence) AddDistributor(d distributor.DTO) error {
	fileAccess := file.New(p.fileName)
	distributors := distributor.Distributors{}
	err := fileAccess.Get(&distributors)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not retrieve saved distributors: %w", err)
	}

	distributors[d.Name()] = d

	err = fileAccess.Save(distributors)
	if err != nil {
		return fmt.Errorf("could not save updated distributors list: %w", err)
	}
	return nil
}

func (p persistence) GetDistributors() (distributor.Distributors, error) {
	fileAccess := file.New(p.fileName)
	distributors := distributor.Distributors{}
	err := fileAccess.Get(&distributors)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("could not get saved distributors: %w", err)
	}
	return distributors, nil
}

func (p persistence) DeleteDistributor(name distributor.Name) error {
	fileAccess := file.New(p.fileName)
	distributors := distributor.Distributors{}
	err := fileAccess.Get(&distributors)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could not get saved distributors: %w", err)
	}

	delete(distributors, name)
	err = fileAccess.Save(distributors)
	if err != nil && err != io.EOF {
		return fmt.Errorf("could save distributors: %w", err)
	}
	return nil
}

func New() distributor.Persistence {
	return persistence{fileName: fileName}
}

const fileName = "distribuidor.ob"
