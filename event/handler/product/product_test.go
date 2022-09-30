package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"obra-blanca/event"
	"obra-blanca/event/handler/dto"
	"obra-blanca/event/product"
	"testing"
)

func TestProductHandler_POST_FailsOnEmptyBody(t *testing.T) {
	var h http.Handler
	h, err := GetManagerHandler()
	if err != nil {
		t.Fatal(err.Error())
	}
	responseRecorder := httptest.ResponseRecorder{}

	h.ServeHTTP(&responseRecorder, httptest.NewRequest(http.MethodPost, "http://localhost:8080", http.NoBody))
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("[POST] did not fail on bad request (empty body): %d", responseRecorder.Code)
	}
}

func TestProductHandler_POST(t *testing.T) {
	var h http.Handler = managerHandler{persistence: fakePersistence{}}
	responseRecorder := httptest.ResponseRecorder{}

	p := product.DTO{
		ProductName:      "test-product",
		ProductInventory: product.InventoryDTO{InitialAmount: 0, Aside: 0},
		ProductUnit:      "test-unit",
	}

	productJson, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := bytes.NewReader(productJson)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/director/product?event="+fakeEventName, body)

	h.ServeHTTP(&responseRecorder, request)
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("[POST] failed on product post: %d", responseRecorder.Code)
	}
}

type fakePersistence struct {
}

func (f fakePersistence) SaveEventDto(_ *event.DTO) error {
	return nil
}

func (f fakePersistence) GetEventDTOs() (dto.DTOs, error) {
	dtos := make(dto.DTOs)
	dtos[fakeEventName] = event.DTO{EventCatalog: map[product.Name]product.DTO{}}
	return dtos, nil
}

const fakeEventName = "some-event"
