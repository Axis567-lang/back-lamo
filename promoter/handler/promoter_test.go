package handler

import (
	"net/http"
	"testing"
)

func TestGetHandler(t *testing.T) {
	var _ http.Handler = GetHandler()
}
