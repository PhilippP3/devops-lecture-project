package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PhilippP3/devops-lecture-project/checkout-service/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler).Methods("POST")
	port := 8082
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// test
