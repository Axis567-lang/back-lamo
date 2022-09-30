package session

import (
	"fmt"
	"obra-blanca/login/token"
	"obra-blanca/login/user"
)

func NewValidator() (Validator, error) {
	p, err := NewPersistence(defaultFileName)
	if err != nil {
		return nil, err
	}
	return validator{persistence: p}, nil
}

type Validator interface {
	Validate(token token.Token) (user.User, error)
}

type validator struct {
	persistence Persistence
}

func (v validator) Validate(token token.Token) (user.User, error) {
	session, err := v.persistence.GetSession(token)
	if err != nil {
		return user.User{}, fmt.Errorf("could not get session: %w", err)
	}

	return session.User, nil
}
