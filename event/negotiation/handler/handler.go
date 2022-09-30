package handler

import (
	"encoding/json"
	"net/http"
	"obra-blanca/distributor"
	"obra-blanca/event"
	eventHandler "obra-blanca/event/handler"
	"obra-blanca/event/negotiation"
	"obra-blanca/event/product"
)

type handler struct {
}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getNegotiations(writer, request)
	case http.MethodPost:
		createNegotiation(writer, request)
	case http.MethodPut:
		addProductToNegotiation(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func addProductToNegotiation(writer http.ResponseWriter, request *http.Request) {
	eventName := request.FormValue("event")
	if eventName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty event name"))
		return
	}
	eventPersistence, err := eventHandler.GetPersistence()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get event persistence: " + err.Error()))
		return
	}
	events, err := eventPersistence.GetEventDTOs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get events from persistence: " + err.Error()))
		return
	}

	e := events[event.Name(eventName)]
	distributorName := request.FormValue("distributor")
	if distributorName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("cannot get events from persistence: " + err.Error()))
		return
	}
	negotiationDTO := e.Assignments()[distributor.Name(distributorName)]

	decoder := json.NewDecoder(request.Body)
	var p product.DTO
	err = decoder.Decode(&p)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("cannot parse product from persistence: " + err.Error()))
		return
	}
	negotiationDTO.AddProduct(p, p.Name())
	e.EventAssignments[distributor.Name(distributorName)] = negotiationDTO
	err = eventPersistence.SaveEventDto(&e)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot save product in negotiation to persistence: " + err.Error()))
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func createNegotiation(writer http.ResponseWriter, request *http.Request) {
	eventName := request.FormValue("event")
	if eventName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty event name"))
		return
	}
	distributorName := request.FormValue("distributor")
	if distributorName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty distributor name"))
		return
	}
	eventPersistence, err := eventHandler.GetPersistence()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get event persistence: " + err.Error()))
		return
	}
	events, err := eventPersistence.GetEventDTOs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get events from persistence: " + err.Error()))
		return
	}

	e := events[event.Name(eventName)]

	if e.EventAssignments == nil {
		e.EventAssignments = make(map[distributor.Name]negotiation.DTO)
	}
	e.EventAssignments[distributor.Name(distributorName)] = *negotiation.New()
	err = eventPersistence.SaveEventDto(&e)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get events from persistence: " + err.Error()))
		return
	}
}

func getNegotiations(writer http.ResponseWriter, request *http.Request) {
	eventName := request.FormValue("event")
	if eventName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty event name"))
		return
	}
	eventPersistence, err := eventHandler.GetPersistence()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get event persistence: " + err.Error()))
		return
	}
	events, err := eventPersistence.GetEventDTOs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot get events from persistence: " + err.Error()))
		return
	}

	e := events[event.Name(eventName)]
	negotiations := make(map[distributor.Name]negotiation.DTO)
	for name, dto := range e.EventAssignments {
		negotiations[name] = dto
	}

	distributorName := request.FormValue("distributor")
	if distributorName != "" {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(negotiations[distributor.Name(distributorName)])
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte("cannot encode negotiation to API: " + err.Error()))
			return
		}
		return
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(negotiations)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("cannot encode negotiations to API: " + err.Error()))
		return
	}
}

func GetHandler() http.Handler {
	return handler{}
}
