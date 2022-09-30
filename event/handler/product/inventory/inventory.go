package inventory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"obra-blanca/event"
	eventHandler "obra-blanca/event/handler"
	"obra-blanca/event/product"
	"strconv"
)

func GetHandler(handlerType HandlerType) (http.Handler, error) {
	switch handlerType {
	case Manager:
		return managerHandler{}, nil
	case Promoter:
		return promoterHandler{}, nil
	default:
		return nil, fmt.Errorf("invalid inventory handler type: %s", handlerType)
	}

}

type managerHandler struct {
}

func (h managerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		logInventory(writer, request)
	case http.MethodGet:
		readInventory(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func readInventory(writer http.ResponseWriter, request *http.Request) {
	eventName := request.FormValue(eventKey)
	productName := request.FormValue(productKey)

	if eventName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty event id"))
		return
	}
	if productName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty product id"))
		return
	}
	eventPersistence, err := eventHandler.GetPersistence()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not retrieve event persistence: %s", err.Error())))
		return
	}
	events, err := eventPersistence.GetEventDTOs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not get events from persistence: %s", err.Error())))
		return
	}
	if _, exists := events[event.Name(eventName)]; !exists {
		writer.WriteHeader(http.StatusFailedDependency)
		_, _ = writer.Write([]byte(fmt.Sprintf("event not found: %s", eventName)))
		return
	}

	eventDto := events[event.Name(eventName)]
	if _, exists := eventDto.EventCatalog[product.Name(productName)]; !exists {
		writer.WriteHeader(http.StatusFailedDependency)
		_, _ = writer.Write([]byte(fmt.Sprintf("product not found: %s", productName)))
		return
	}
	p := eventDto.EventCatalog[product.Name(productName)]
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(p.ProductInventory)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not save product: %s", err.Error())))
		return
	}
}

func logInventory(writer http.ResponseWriter, request *http.Request) {
	eventName := request.FormValue(eventKey)
	productName := request.FormValue(productKey)
	amountString := request.FormValue(amountKey)

	if eventName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty event id"))
		return
	}
	if productName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty product id"))
		return
	}
	if amountString == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("empty amountString"))
		return
	}
	amt, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not parse amount to float: %s", err.Error())))
		return
	}

	eventPersistence, err := eventHandler.GetPersistence()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not retrieve event persistence: %s", err.Error())))
		return
	}
	events, err := eventPersistence.GetEventDTOs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not get events from persistence: %s", err.Error())))
		return
	}

	if _, exists := events[event.Name(eventName)]; !exists {
		writer.WriteHeader(http.StatusFailedDependency)
		_, _ = writer.Write([]byte(fmt.Sprintf("event not found: %s", eventName)))
		return
	}
	eventDto := events[event.Name(eventName)]

	if _, exists := eventDto.EventCatalog[product.Name(productName)]; !exists {
		writer.WriteHeader(http.StatusFailedDependency)
		_, _ = writer.Write([]byte(fmt.Sprintf("product not found: %s", productName)))
		return
	}
	p := eventDto.EventCatalog[product.Name(productName)]

	i := eventDto.EventCatalog[product.Name(productName)].ProductInventory

	i.InitialAmount = amt
	p.ProductInventory = i
	eventDto.EventCatalog[product.Name(productName)] = p

	err = eventPersistence.SaveEventDto(&eventDto)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("could not save DTO: %s", err.Error())))
		return
	}
	writer.WriteHeader(200)
}

type HandlerType string

const Manager HandlerType = "manager"
const Promoter HandlerType = "promoter"

const eventKey = "event"
const productKey = "product"
const amountKey = "amount"
