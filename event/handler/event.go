package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogelioConsejo/golibs/file"
	"net/http"
	"obra-blanca/event"
	"obra-blanca/event/handler/persistence"
	"strings"
	"time"
)

func NewEventCreatorHandler() (CreatorHandler, error) {
	p, err := GetPersistence()
	if err != nil {
		return CreatorHandler{}, err
	}
	return CreatorHandler{persistence: p}, nil
}

type CreatorHandler struct {
	persistence Persistence
	writer      http.ResponseWriter
}

func (c CreatorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", allowedCORSOrigins)

	c.writer = writer
	switch request.Method {
	case http.MethodPost:
		c.createEvent(writer, request)
	case http.MethodGet:
		c.getEvents(writer, request)
	case http.MethodOptions:
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		respond(writer, http.StatusNoContent, "")
	default:
		respond(writer, http.StatusMethodNotAllowed, fmt.Sprintf("Method not allowed: %s\n", request.Method))
	}
}

func (c CreatorHandler) getEvents(w http.ResponseWriter, r *http.Request) {
	DTOs, err := c.persistence.GetEventDTOs()
	if err != nil {
		respond(w, http.StatusInternalServerError, fmt.Sprintf("error getting events: %s", err.Error()))
		return
	}

	eventId := r.FormValue(eventIdKey)
	if eventId == "" {
		dtosJSON, err := json.Marshal(DTOs)
		if err != nil {
			respond(w, http.StatusBadRequest, fmt.Sprintf("error getting events: %s", err.Error()))
			return
		}
		respond(w, http.StatusOK, string(dtosJSON))
		return
	}

	if DTO, exists := DTOs[event.Name(eventId)]; exists {
		dtoJSON, err := json.Marshal(DTO)
		if err != nil {
			respond(w, http.StatusBadRequest, fmt.Sprintf("error getting events: %s", err.Error()))
			return
		}
		respond(w, http.StatusOK, string(dtoJSON))
		return
	}

	respond(w, http.StatusNotFound, fmt.Sprintf("event not found: %s", eventId))
	return
}

func (c CreatorHandler) createEvent(w http.ResponseWriter, r *http.Request) {
	c.writer = w

	e, err := getEventDtoFromRequest(r)
	if err != nil {
		respond(w, http.StatusBadRequest, fmt.Sprintf("error getting event from request: %s", err.Error()))
		return
	}

	if err := validateEventRequest(w, e); err != nil {
		respond(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err.Error()))
		return
	}
	if err := c.persistence.SaveEventDto(e); err != nil {
		respond(w, http.StatusInternalServerError, fmt.Sprintf("error saving event: %s", err.Error()))
		return
	}
	respond(w, http.StatusCreated, "")
	return

}

func respond(writer http.ResponseWriter, statusCode int, msg string) {
	writer.WriteHeader(statusCode)
	_, _ = writer.Write([]byte(msg))
}

func GetPersistence() (Persistence, error) {
	p, err := persistence.New(UseDefault)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func validateEventRequest(w http.ResponseWriter, e *event.DTO) error {
	if !assertNameNotEmpty(w, e) {
		return fmt.Errorf("empty name")
	}
	if !assertTimeNotEmpty(w, e) {
		return fmt.Errorf("empty date")
	}
	return nil
}

func assertTimeNotEmpty(w http.ResponseWriter, e *event.DTO) bool {
	if e.EventTime.Equal(time.Time{}) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty event starting date"))
		return false
	}
	return true
}

func assertNameNotEmpty(w http.ResponseWriter, e *event.DTO) bool {
	if strings.TrimSpace(string(e.EventName)) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("empty name"))
		return false
	}
	return true
}

func getEventDtoFromRequest(r *http.Request) (*event.DTO, error) {
	if r.Body == nil || r.Body == http.NoBody {
		return nil, EmptyRequestBody
	}
	decoder := json.NewDecoder(r.Body)
	e := &event.DTO{}
	err := decoder.Decode(e)
	return e, err
}

var EmptyRequestBody error = errors.New("empty request body")

const UseDefault file.Name = ""
const allowedCORSOrigins = "*"
const eventIdKey = "id"
