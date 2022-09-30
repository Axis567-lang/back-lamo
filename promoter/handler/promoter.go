package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"obra-blanca/promoter"
	"obra-blanca/promoter/persistence"
	"strings"
)

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerNewPromoter(w, r)
	case http.MethodGet:
		getPromoters(w, r)
	case http.MethodDelete:
		deletePromoter(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func deletePromoter(w http.ResponseWriter, r *http.Request) {
	promoterName := r.FormValue("name")
	if promoterName == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty promoter name"))
		return
	}
	var p promoter.Persistence = persistence.New()
	err := p.DeletePromoter(promoter.Name(promoterName))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not delete promoter: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getPromoters(w http.ResponseWriter, _ *http.Request) {
	var p promoter.Persistence = persistence.New()
	promoters, err := p.GetPromoters()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not read promoters dtos from persistence: %s", err.Error())))
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(promoters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not write promoters dtos to API: %s", err.Error())))
		return
	}
}

func registerNewPromoter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto promoter.DTO
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("could not decode promoter dto from request: %s", err.Error())))
		return
	}

	if strings.TrimSpace(string(dto.Name())) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty promoter name"))
		return
	}

	var p promoter.Persistence = persistence.New()
	err = p.AddPromoter(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not add promoter to persistence: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetHandler() http.Handler {
	return handler{}
}
