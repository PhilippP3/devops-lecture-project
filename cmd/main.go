package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/PhilippP3/devops-lecture-project/pkg/auth"
	"github.com/PhilippP3/devops-lecture-project/pkg/products"
	"github.com/PhilippP3/devops-lecture-project/internal/handler"
	"github.com/gorilla/mux"
)

var secretKey = []byte("secret-key")
auth.

func main() {
	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", authLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", authLogoutHandler).Methods("POST")
	// Product Service
	router.HandleFunc("/products", productListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productDetailHandler).Methods("GET")
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler).Methods("POST")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

