package session

import (
	"github.com/rogelioConsejo/golibs/file"
	"obra-blanca/login/token"
	"obra-blanca/login/user"
	"os"
	"testing"
	"time"
)

func TestSession_ShouldPersist(t *testing.T) {
	clean(t, testFileName)
	p, _ := NewPersistence(testFileName)
	per := p.(*persistence)
	newSession := per.makeTestSession(testUser)
	tok := newSession.GetToken()

	var p2, _ = NewPersistence(testFileName)
	retrievedSession, err := p2.GetSession(tok)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	newSessionToken := newSession.GetToken().String()
	retrievedSessionToken := retrievedSession.GetToken().String()
	if newSessionToken != retrievedSessionToken {
		t.Fatalf("did not retrieve the session correctly")
	}

	clean(t, testFileName)
}

func TestSession_Token(t *testing.T) {
	clean(t, testFileName)
	per, _ := NewPersistence(testFileName)
	p := per.(*persistence)
	testSession := p.makeTestSession(testUser)
	var testToken = testSession.GetToken()
	if tokenSize := len(testToken.String()); tokenSize != token.Size {
		t.Fatalf("invalid token size (%d): %s\n", tokenSize, testToken.String())
	}
	clean(t, testFileName)
}

func TestSession_Token_ShouldNotChange(t *testing.T) {
	clean(t, testFileName)
	per, _ := NewPersistence(testFileName)
	p := per.(*persistence)
	testSession := p.makeTestSession(testUser)
	if testSession.GetToken().String() != testSession.GetToken().String() {
		t.Fatalf("token should not change")
	}
	clean(t, testFileName)
}

func (p *persistence) makeTestSession(usr user.User) Session {
	var expirationDate = time.Time{}
	var testSession, _ = p.New(usr, expirationDate)
	return testSession
}

var testUser = user.User{
	Username: "test-user",
	Password: "test-password",
}

const testFileName = "test-persistence.ob"

func clean(t *testing.T, fileName file.Name) {
	_ = os.Remove(string(fileName))
}
