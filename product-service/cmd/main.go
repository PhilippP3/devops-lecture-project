package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PhilippP3/devops-lecture-project/product-service/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Product Service
	router.HandleFunc("/products", handler.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handler.ProductDetailHandler).Methods("GET")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
