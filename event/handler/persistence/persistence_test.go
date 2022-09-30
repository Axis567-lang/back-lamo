package persistence

import (
	"fmt"
	"obra-blanca/event"
	"os"
	"testing"
	"time"
)

func TestFilePersistence_SaveEventDto(t *testing.T) {
	const testFilename = "eventos.ob.test"
	filePersistence, err := New(testFilename)
	if err != nil {
		t.Fatalf("could not create fileName persistence access: %s", err.Error())
	}
	err = filePersistence.SaveEventDto(&event.DTO{EventName: expectedEventName, EventTime: time.Now()})
	if err != nil {
		t.Fatalf("could not save event DTO: %s", err.Error())
	}
	secondFilePersistenceAccess, err := New(testFilename)
	if err != nil {
		t.Fatalf("could not create second fileName persistence access: %s", err.Error())
	}

	dtos, err := secondFilePersistenceAccess.GetEventDTOs()
	if err != nil {
		t.Fatalf("could not get events DTOs: %s", err.Error())
	}
	if dtos == nil {
		t.Fatalf("nil map retrieved from persistence")
	}
	if _, exists := dtos[expectedEventName]; !exists {
		t.Fatalf("event '%s' not found", expectedEventName)
	}

	err = os.Remove(testFilename)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}

const expectedEventName = "expected-event"
