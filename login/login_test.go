package login

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"obra-blanca/login/session"
	"obra-blanca/login/token"
	"obra-blanca/login/user"
	"strings"
	"testing"
	"time"
)

func TestLoginHandler_ServeHTTP_ShouldReturnToken(t *testing.T) {
	w, r := simulateLoginWebRequest(http.MethodGet)
	handler{authenticationVerifier: stubLoginPersistence{}, persistence: stubLoginPersistence{}}.ServeHTTP(&w, r)
	if w.statusCode != 200 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, "token") {
		t.Fatalf("token not found: %s", w.body)
	}
}

func TestLoginHandler_ServeHTTP_ShouldFailOnInvalidKeys(t *testing.T) {
	w := spyWriter{}
	r := httptest.NewRequest(http.MethodGet, "localhost:8080", http.NoBody)
	r.SetBasicAuth("invalid", "not valid")
	handler{authenticationVerifier: stubLoginPersistence{}}.ServeHTTP(&w, r)
	if w.statusCode != 401 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, badCredentialsMessage) {
		t.Fatalf("invalid error message: %s", w.body)
	}
}

func TestLoginHandler_ServeHTTP_ShouldFailWithoutBasicAuth(t *testing.T) {
	w, r := simulateLoginWebRequest(http.MethodGet, true)
	handler{}.ServeHTTP(&w, r)
	if w.statusCode != 400 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, missingAuthErrorMessage) {
		t.Fatalf("invalid error message: %s", w.body)
	}
}

func TestLoginHandler_ServeHTTP_MethodShouldNotBePost(t *testing.T) {
	w, r := simulateLoginWebRequest(http.MethodPost)
	handler{}.ServeHTTP(&w, r)
	if w.statusCode != 405 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, methodErrorMessage) {
		t.Fatalf("invalid error message: %s", w.body)
	}
}
func TestLoginHandler_ServeHTTP_MethodShouldNotBePut(t *testing.T) {
	w, r := simulateLoginWebRequest(http.MethodPut)
	handler{}.ServeHTTP(&w, r)
	if w.statusCode != 405 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, methodErrorMessage) {
		t.Fatalf("invalid error message: %s", w.body)
	}
}
func TestLoginHandler_ServeHTTP_MethodShouldNotBeDelete(t *testing.T) {
	w, r := simulateLoginWebRequest(http.MethodDelete)
	handler{}.ServeHTTP(&w, r)
	if w.statusCode != 405 {
		t.Fatalf(fmt.Sprintf("%d", w.statusCode))
	}
	if !strings.Contains(w.body, methodErrorMessage) {
		t.Fatalf("invalid error message: %s", w.body)
	}
}

func simulateLoginWebRequest(httpMethod string, noAuth ...bool) (spyWriter, *http.Request) {
	w := spyWriter{}
	r := httptest.NewRequest(httpMethod, "localhost:8080", http.NoBody)
	if len(noAuth) == 0 {
		r.SetBasicAuth(string(expectedUserName), string(expectedPassword))
	}
	return w, r
}

type stubLoginPersistence struct {
}

func (s stubLoginPersistence) New(u user.User, time time.Time) (session.DTO, error) {
	return session.DTO{}, nil
}

func (s stubLoginPersistence) GetSession(token token.Token) (session.DTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s stubLoginPersistence) CheckUser(u user.Username, p user.Password) (bool, error) {
	expectedUserNameHash := expectedUserName
	expectedPasswordHash := expectedPassword
	if u == expectedUserNameHash && p == expectedPasswordHash {
		return true, nil
	}
	return false, nil
}

type spyWriter struct {
	statusCode int
	body       string
}

func (s *spyWriter) Header() http.Header {
	//TODO implement me
	panic("implement me")
}

func (s *spyWriter) Write(bytes []byte) (int, error) {
	s.body = s.body + string(bytes)
	return len(bytes), nil
}

func (s *spyWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
}

const expectedUserName user.Username = "usr"
const expectedPassword user.Password = "pass"
