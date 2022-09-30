package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"obra-blanca/distributor"
	"obra-blanca/distributor/persistence"
	"strings"
)

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerNewDistributor(w, r)
	case http.MethodGet:
		getDistributors(w, r)
	case http.MethodDelete:
		deleteDistributor(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func deleteDistributor(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty promoter name"))
		return
	}
	var p distributor.Persistence = persistence.New()
	err := p.DeleteDistributor(distributor.Name(name))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not delete promoter: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getDistributors(w http.ResponseWriter, _ *http.Request) {
	var p distributor.Persistence = persistence.New()
	dtos, err := p.GetDistributors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not read distributor dtos from persistence: %s", err.Error())))
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(dtos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("could not write distributor dtos to API: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func registerNewDistributor(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto distributor.DTO
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("could not decode distributor dto from request: %s", err.Error())))
		return
	}

	if strings.TrimSpace(string(dto.Name())) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty distributor name"))
		return
	}

	var p distributor.Persistence = persistence.New()
	err = p.AddDistributor(dto)
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
