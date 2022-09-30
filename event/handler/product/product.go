package product

import (
	"encoding/json"
	"net/http"
	"obra-blanca/event"
	eventHandler "obra-blanca/event/handler"
	"obra-blanca/event/handler/persistence"
	"obra-blanca/event/product"
)

func GetManagerHandler() (http.Handler, error) {
	p, err := persistence.New("")
	if err != nil {
		return nil, err
	}
	return managerHandler{persistence: &p}, nil
}

func GetPromoterHandler() (http.Handler, error) {
	p, err := persistence.New("")
	if err != nil {
		return nil, err
	}
	return promoterHandler{managerHandler{persistence: &p}}, nil
}

type promoterHandler struct {
	handler managerHandler
}

func (p promoterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		p.handler.getEventProducts(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type managerHandler struct {
	persistence eventHandler.Persistence
}

func (h managerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.registerProduct(w, r)
	case http.MethodGet:
		h.getEventProducts(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h managerHandler) registerProduct(w http.ResponseWriter, r *http.Request) {
	if checkForEmptyBody(w, r) {
		return
	}

	eventId, done := getEventId(w, r)
	if done {
		return
	}

	var dto product.DTO
	dto, done = getDTOFromRequest(w, r)
	if done {
		return
	}

	ev, done := getEvent(w, h, eventId)
	if done {
		return
	}

	if ev.EventCatalog == nil {
		ev.EventCatalog = make(event.Catalog)
	}
	ev.EventCatalog[dto.ProductName] = dto
	err := h.persistence.SaveEventDto(&ev)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}

func (h managerHandler) getEventProducts(w http.ResponseWriter, r *http.Request) {
	eventId, done := getEventId(w, r)
	if done {
		return
	}

	ev, done := getEvent(w, h, eventId)
	if done {
		return
	}
	products := ev.Catalog
	encoder := json.NewEncoder(w)
	err := encoder.Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

func getEvent(w http.ResponseWriter, h managerHandler, eventId string) (ev event.DTO, done bool) {
	eventDTOs, err := h.persistence.GetEventDTOs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return event.DTO{}, true
	}
	ev = eventDTOs[event.Name(eventId)]
	return ev, false
}

func getDTOFromRequest(w http.ResponseWriter, r *http.Request) (dto product.DTO, done bool) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return product.DTO{}, true
	}
	return dto, false
}

func getEventId(w http.ResponseWriter, r *http.Request) (id string, done bool) {
	eventId := r.FormValue("event")
	if eventId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(emptyEventErrorMessage))
		return "", true
	}
	return eventId, false
}

func checkForEmptyBody(w http.ResponseWriter, r *http.Request) bool {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(emptyBodyErrorMessage))
		return true
	}
	return false
}

const emptyBodyErrorMessage = "body is empty"
const emptyEventErrorMessage = "event was not specified"
