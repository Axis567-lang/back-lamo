package login

import (
	"obra-blanca/login/persistence"
	"obra-blanca/login/user"
)

func NewValidator() Validator {
	return validator{}
}

type Validator interface {
	ValidateUser(usr user.Username, pass user.Password) (bool, error)
}

type validator struct {
}

func (v validator) ValidateUser(usr user.Username, pass user.Password) (bool, error) {
	p, err := persistence.New("")
	if err != nil {
		return false, err
	}
	isValid, err := p.CheckUser(usr, pass)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
