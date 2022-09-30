package authorization

import (
	"net/http"
	"obra-blanca/login/session"
	"obra-blanca/login/token"
	"strings"
)

func SecureHandler(h http.Handler) (http.Handler, error) {
	v, err := session.NewValidator()
	if err != nil {
		return nil, err
	}
	return secureHandler{
		handler:   h,
		validator: v,
	}, nil
}

type secureHandler struct {
	handler   http.Handler
	validator session.Validator
}

func (s secureHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	authHeader := request.Header.Get(HeaderKey)
	authFields := strings.Fields(authHeader)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, access-control-allow-origin, access-control-allow-headers")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

	if request.Method == "OPTIONS"	{
		if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(NoAccessTokenFoundErrorMessage))
			return
		}
	
		t := authFields[1]
		_, err := s.validator.Validate(token.Token{Value: t})
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			_, _ = writer.Write([]byte(InvalidAccessTokenErrorMessage))
			return
		}
		
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	s.handler.ServeHTTP(writer, request)
}

const HeaderKey = "Authorization"
const NoAccessTokenFoundErrorMessage = "no access token found"
const InvalidAccessTokenErrorMessage = "invalid access token"
