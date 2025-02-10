package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PhilippP3/devops-lecture-project/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", handler.AuthLogoutHandler).Methods("POST")
	// Product Service
	router.HandleFunc("/products", handler.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handler.ProductDetailHandler).Methods("GET")
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler).Methods("POST")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
