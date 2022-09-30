package event

import (
	"obra-blanca/distributor"
	"testing"
	"time"
)

func TestEvent_Catalog(t *testing.T) {
	testEvent := createEvent()
	var _ Catalog = testEvent.Catalog()
}

func TestEvent_Assignments(t *testing.T) {
	testEvent := createEvent()
	expectedDistributor := fakeDistributor{}
	testEvent.Assign(expectedDistributor).To(testEvent)
	var testAssignments Assignments = testEvent.Assignments()
	_, exists := testAssignments[expectedDistributor.Name()]
	if !exists {
		t.Fatal("distributor not assigned to promoter for test event")
	}

}

func TestEvent_Name(t *testing.T) {
	testEvent := createEvent()
	var testName Name = testEvent.Name()
	if testName != testEventName {
		t.Fatalf("did not get expected name (%s): got %s\n", testEventName, testName)
	}
}

func TestEvent_Time(t *testing.T) {
	testEvent := createEvent()
	time.Sleep(1000)
	var testTime time.Time = testEvent.Time()
	if testTime.Unix() != now.Unix() {
		t.Fatalf("did not get expected time (%s): got %s\n", now.String(), testTime.String())
	}
}

func TestEvent_AddProduct(t *testing.T) {
	testEvent := createEvent()
	var testDistributor distributor.Distributor = fakeDistributor{}
	testEvent.Assign(testDistributor).To(testEvent)
}

func createEvent() DTO {
	var testEvent DTO = New(now, testEventName)
	return testEvent
}

var _ Event = &DTO{}

const testEventName Name = "test-event"

var now = time.Now()
