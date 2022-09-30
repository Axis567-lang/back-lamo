package authorization

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"obra-blanca/login/token"
	"obra-blanca/login/user"
	"strings"
	"testing"
)

func TestSecureHandler_ShouldRequireAToken(t *testing.T) {
	var handler http.Handler = stubHandler{}
	var s http.Handler
	s, err := SecureHandler(handler)
	securedHandler := s.(secureHandler)
	securedHandler.validator = fakeValidator{}
	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.ResponseRecorder{}
	r := httptest.NewRequest(http.MethodGet, "localhost:8080", http.NoBody)
	securedHandler.ServeHTTP(&w, r)
	b, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if w.Code != http.StatusBadRequest || !strings.Contains(NoAccessTokenFoundErrorMessage, fmt.Sprintf("%s", b)) {
		t.Fatalf("secured handler did not fail on missing token: %s", w.Body.String())
	}

	r.Header.Set(HeaderKey, "bearer "+expectedToken)
	w2 := httptest.ResponseRecorder{}
	securedHandler.ServeHTTP(&w2, r)
	if w2.Code != http.StatusOK {
		t.Fatalf("secured handler failed to recognize token")
	}

}

type stubHandler struct {
}

func (f stubHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type fakeValidator struct {
}

func (f fakeValidator) Validate(token token.Token) (user.User, error) {
	if token.String() == expectedToken {
		return user.User{Username: expectedUser, Password: expectedPass}, nil
	}
	return user.User{}, errors.New("user not found")
}

const expectedUser = "expected-user"
const expectedPass = "expected-pass"
const expectedToken = "expected-token"
