package login

import (
	"encoding/json"
	"net/http"
	"obra-blanca/login/persistence"
	"obra-blanca/login/session"
	"obra-blanca/login/user"
	"time"
)

func NewLoginHandler() (handler, error) {
	v, err := persistence.New("")
	if err != nil {
		return handler{}, err
	}
	p, err := session.NewPersistence("")
	if err != nil {
		return handler{}, err
	}
	return handler{authenticationVerifier: v, persistence: p}, nil
}

type handler struct {
	authenticationVerifier user.AuthenticationVerifier
	persistence            session.Persistence
}

func (l handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, access-control-allow-origin, access-control-allow-headers")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	switch request.Method {
	case http.MethodGet:
		usr, pass, ok := request.BasicAuth()
		if !ok {
			writer.WriteHeader(400)
			_, _ = writer.Write([]byte(missingAuthErrorMessage))
			return
		}

		authenticated, err := l.authenticationVerifier.CheckUser(user.Username(usr), user.Password(pass))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(authenticationErrorMessage + err.Error()))
			return
		}
		if !authenticated {
			writer.WriteHeader(401)
			_, _ = writer.Write([]byte(badCredentialsMessage))
			return
		}

		dto, err := l.persistence.New(user.User{
			Username: user.Username(usr),
			Password: user.Password(pass),
		}, time.Now().Add(2*time.Hour))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte("could not create session"))
			return
		}

		encoder := json.NewEncoder(writer)
		err = encoder.Encode(dto)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte("could not save session"))
			return
		}

		writer.WriteHeader(http.StatusOK)
	case http.MethodOptions:
		writer.WriteHeader(http.StatusOK)
		return
	default:
		writer.WriteHeader(405)
		_, _ = writer.Write([]byte(methodErrorMessage))
		return
	}

}

const methodErrorMessage = "http method should be GET"
const missingAuthErrorMessage = "request does not include basic auth"
const authenticationErrorMessage = "could not authenticate user: "
const badCredentialsMessage = "invalid credentials"
