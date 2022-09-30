package inventory

import "net/http"

type promoterHandler struct {
}

func (p promoterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		readInventory(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
