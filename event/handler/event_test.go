package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"obra-blanca/event"
	"obra-blanca/event/handler/dto"
	"strings"
	"testing"
)

func TestNewEventCreatorHandler(t *testing.T) {
	_, _ = NewEventCreatorHandler()
}

func TestEventCreatorHandler_ServeHTTP_FailsIfUsingPut(t *testing.T) {
	var p Persistence = &fakePersistence{}
	ec := makeEventCreator(p)
	w, r := simulateEventWebRequest(http.MethodPut, expectedEventName, expectedTime)
	ec.ServeHTTP(w, r)
	if w.statusCode >= 200 && w.statusCode <= 299 {
		t.Fatalf("did not detect invalid PUT request")
	}
}

// TODO
func TestEventCreatorHandler_GetEvents(t *testing.T) {
}

func TestEventCreatorHandler_CreateEvent_SavesEvent(t *testing.T) {
	var testPersistence Persistence = &fakePersistence{}
	ec := makeEventCreator(testPersistence)
	w, r := simulateEventWebRequest(http.MethodPost, expectedEventName, expectedTime)
	ec.createEvent(w, r)
	if w.statusCode < 200 || w.statusCode > 299 {
		t.Fatalf("error creating event: %d - %s", w.statusCode, w.body)
	}

	retrievedEvent, err := testPersistence.(*fakePersistence).getEvent(expectedEventName)
	if err != nil {
		t.Fatalf("could not retrieve event: %s\n", err.Error())
	}
	if retrievedEventName := retrievedEvent.Name(); retrievedEventName != expectedEventName {
		t.Fatalf("invalid event name: %s -- expected: %s\n", retrievedEventName, expectedEventName)
	}
}

func TestEventCreatorHandler_CreateEvent_FailsOnEmptyBody(t *testing.T) {
	var p Persistence = &fakePersistence{}
	ec := makeEventCreator(p)
	w, r := simulateEventWebRequest(http.MethodPost, expectedEventName, expectedTime)
	r.Body = http.NoBody
	ec.createEvent(w, r)
	if w.statusCode >= 200 && w.statusCode <= 299 {
		t.Fatalf("did not detect invalid empty body")
	}
}

func TestEventCreatorHandler_CreateEvent_FailsOnEmptyName(t *testing.T) {
	var p Persistence = &fakePersistence{}
	ec := makeEventCreator(p)
	w, r := simulateEventWebRequest(http.MethodPost, "", expectedTime)
	ec.createEvent(w, r)
	if w.statusCode >= 200 && w.statusCode <= 299 {
		t.Fatalf("did not detect invalid empty event name")
	}
}

func TestEventCreatorHandler_CreateEvent_FailsOnEmptyTime(t *testing.T) {
	var p Persistence = &fakePersistence{}
	ec := makeEventCreator(p)
	w, r := simulateEventWebRequest(http.MethodPost, expectedEventName, "")
	ec.createEvent(w, r)
	if w.statusCode >= 200 && w.statusCode <= 299 {
		t.Fatalf("did not detect invalid empty event name")
	}
}

func TestEventCreatorHandler_CreateEvent_FailsOnErrorWhileSaving(t *testing.T) {
	var p Persistence = &stubErrorSavingPersistence{}
	ec := makeEventCreator(p)
	w, r := simulateEventWebRequest(http.MethodPost, expectedEventName, expectedTime)
	ec.createEvent(w, r)
	if w.statusCode >= 200 && w.statusCode <= 299 {
		t.Fatalf("did not detect invalid empty event name")
	}
}

type stubErrorSavingPersistence struct {
}

func (p stubErrorSavingPersistence) GetEventDTOs() (dto.DTOs, error) {
	//TODO implement me
	panic("implement me")
}

func (p stubErrorSavingPersistence) SaveEventDto(_ *event.DTO) error {
	return errors.New("some error")
}

func simulateEventWebRequest(httpMethod string, name event.Name, time string) (*spyWriter, *http.Request) {
	w := &spyWriter{}
	r := httptest.NewRequest(httpMethod, "localhost:8080", strings.NewReader(fmt.Sprintf(eventRequestFormat, name, time)))
	return w, r
}

func makeEventCreator(p Persistence) CreatorHandler {
	ec := CreatorHandler{persistence: p}
	return ec
}

type fakePersistence struct {
	events map[event.Name]*event.DTO
}

func (p *fakePersistence) GetEventDTOs() (dto.DTOs, error) {
	//TODO implement me
	panic("implement me")
}

func (p *fakePersistence) SaveEventDto(e *event.DTO) error {
	if p.events == nil {
		p.events = make(map[event.Name]*event.DTO)
	}
	p.events[e.EventName] = e
	return nil
}

func (p *fakePersistence) getEvent(eventName event.Name) (event.Event, error) {
	d, exists := p.events[eventName]
	if !exists {
		return nil, fmt.Errorf("event not found: %s\n", eventName)
	}
	return d, nil
}

type spyWriter struct {
	statusCode int
	body       string
}

func (s *spyWriter) Header() http.Header {
	return http.Header{}
}

func (s *spyWriter) Write(bytes []byte) (int, error) {
	s.body = s.body + string(bytes)
	return len(bytes), nil
}

func (s *spyWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
}

const expectedEventName event.Name = "expected-event"
const expectedTime = "2018-12-10T13:45:00.000Z"
const eventRequestFormat = "{\"name\": \"%s\", \"time\": \"%s\"}"
