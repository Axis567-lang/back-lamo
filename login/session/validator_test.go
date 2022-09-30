package session

import (
	"obra-blanca/login/user"
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	clean(t, testFileName)
	sessionPersistence, err := NewPersistence(testFileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fakeUser := user.User{
		Username: "fake",
		Password: "pass",
	}
	dto, err := sessionPersistence.New(fakeUser, time.Now().Add(2*time.Minute))
	if err != nil {
		return
	}

	va, err := NewValidator()
	v := va.(validator)
	v.persistence = sessionPersistence
	var _ user.User
	_, err = v.Validate(dto.Token)
	if err != nil {
		t.Errorf(err.Error())
	}

	clean(t, testFileName)
}
