package persistence

import (
	"obra-blanca/login/user"
	"strings"
	"testing"
)

func TestPersistence(t *testing.T) {
	const expectedUsername = "expected-user"
	const expectedPassword = "expected-pass"
	p1, err := New(testFileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = p1.AddUser(expectedUsername, expectedPassword)
	if err != nil {
		t.Fatalf(err.Error())
	}
	p1.Close()

	p2, err := New(testFileName)
	userIsValid, err := p2.CheckUser(expectedUsername, expectedPassword)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !userIsValid {
		t.Fatalf("user is not valid")
	}

	p2.deleteUser(expectedUsername)
	p2.Close()
}

func TestPersistence_CheckUser_DetectsNonExistingUser(t *testing.T) {
	const invalidUsername = "a"
	var p, err = New(testFileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	userIsValid, err := p.CheckUser(invalidUsername, "a")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if userIsValid {
		t.Fatalf("user '%s' is actually not valid", invalidUsername)
	}
	p.Close()
}

func TestPersistence_CheckUser_ReturnsErrorOnEmptyHashes(t *testing.T) {
	_, err := persistence{}.CheckUser("", "a")
	if err == nil || !strings.Contains(err.Error(), emptyUserErrorMessage) {
		t.Fatalf("did not detect empty usernameHash")
	}
	_, err = persistence{}.CheckUser("a", "")
	if err == nil || !strings.Contains(err.Error(), emptyPasswordErrorMessage) {
		t.Fatalf("did not detect empty passwordHash")
	}
}

var _ user.AuthenticationVerifier = persistence{}

const testFileName = "credentials.ob.test"
