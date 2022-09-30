package main

import (
	"fmt"
	"net/http"
	"obra-blanca/authorization"
	negotiation "obra-blanca/event/negotiation/handler"
	"obra-blanca/login"
	"time"
)

func main() {
	logOutput("OBRA BLANCA 🏗️")
	logOutput("Initializing Back-End Server...")
	logOutput("Start-up time: " + time.Now().String())
	logOutput("Version: " + version)
	secureEventHandler, done := createSecureEventHandler()
	if done {
		return
	}
	secureDirectorProductHandler, done := createSecureProductHandler(managerHandler)
	if done {
		return
	}
	secureProductPromoterHandler, done := createSecureProductHandler(promoterHandler)
	if done {
		return
	}
	secureInventoryManagerHandler, done := createSecureInventoryHandler(managerHandler)
	if done {
		return
	}
	secureInventoryPromoterHandler, done := createSecureInventoryHandler(promoterHandler)
	if done {
		return
	}
	securePromoterManagerHandler, done := createSecurePromoterManagerHandler()
	if done {
		return
	}
	secureDistributorManagerHandler, done := createSecureDistributorManagerHandler()
	if done {
		return
	}
	secureNegotiationsHandler, done := createSecureNegotiationsHandler()
	if done {
		return
	}
	loginHandler, err := login.NewLoginHandler()
	if err != nil {
		die(err)
		return
	}

	http.Handle("/director/product", secureDirectorProductHandler)
	http.Handle("/director/inventory", secureInventoryManagerHandler)
	http.Handle("/director/event", secureEventHandler)
	http.Handle("/director/promoter", securePromoterManagerHandler)
	http.Handle("/director/distributor", secureDistributorManagerHandler)

	http.Handle("/promoter/product", secureProductPromoterHandler)
	http.Handle("/promoter/inventory", secureInventoryPromoterHandler)
	http.Handle("/promoter/negotiation", secureNegotiationsHandler)

	http.Handle("/login", loginHandler)

	logOutput("👍 Everything seems to be working correctly. ")
	logOutput(fmt.Sprintf("Running on Port %d", defaultPort))
	logOutput("*****************************************")
	logOutput(halQuote(hello))

	err = http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)
	if err != nil {
		die(err)
		return
	}
}

func createSecureNegotiationsHandler() (handler http.Handler, done bool) {
	handler = negotiation.GetHandler()
	secureHandler, err := authorization.SecureHandler(handler)
	if err != nil {
		die(err)
		return nil, true
	}
	return secureHandler, false
}

const defaultPort = 8080
const version = "v0.0.4b"
