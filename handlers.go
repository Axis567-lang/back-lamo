package main

import (
	"net/http"
	"obra-blanca/authorization"
	distributor "obra-blanca/distributor/handler"
	event "obra-blanca/event/handler"
	"obra-blanca/event/handler/product"
	"obra-blanca/event/handler/product/inventory"
	promoter "obra-blanca/promoter/handler"
)

func createSecureDistributorManagerHandler() (handler http.Handler, done bool) {
	handler = distributor.GetHandler()
	secureHandler, err := authorization.SecureHandler(handler)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureHandler, false
}

func createSecurePromoterManagerHandler() (handler http.Handler, done bool) {
	handler = promoter.GetHandler()
	secureHandler, err := authorization.SecureHandler(handler)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureHandler, false
}

func createSecureInventoryHandler(ht handlerType) (handler http.Handler, done bool) {
	var err error
	switch ht {
	case managerHandler:
		handler, err = inventory.GetHandler(inventory.Manager)
		if err != nil {
			die(err)
			return nil, true
		}
	case promoterHandler:
		handler, err = inventory.GetHandler(inventory.Promoter)
		if err != nil {
			die(err)
			return nil, true
		}
	}

	secureHandler, err := authorization.SecureHandler(handler)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureHandler, false
}

func createSecureProductHandler(t handlerType) (secureProductHandler http.Handler, done bool) {
	var productHandler http.Handler
	var err error
	switch t {
	case managerHandler:
		productHandler, err = product.GetManagerHandler()
		if err != nil {
			die(err)
			return nil, true
		}
	case promoterHandler:
		productHandler, err = product.GetPromoterHandler()
		if err != nil {
			die(err)
			return nil, true
		}
	}
	secureProductHandler, err = authorization.SecureHandler(productHandler)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureProductHandler, false
}

func createSecureEventHandler() (http.Handler, bool) {
	eventCreator, err := event.NewEventCreatorHandler()
	if err != nil {
		die(err)
		return nil, true
	}
	secureEventCreator, err := authorization.SecureHandler(eventCreator)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureEventCreator, false
}

func die(err error) {
	logOutput("❌ error: " + err.Error())
	logOutput(halQuote(fail))
	return
}

type handlerType string

const managerHandler handlerType = "manager"
const promoterHandler handlerType = "promoter"
