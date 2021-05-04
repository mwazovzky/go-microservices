package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
	data   *data.Products
	logger *log.Logger
}

// NewProducts returns a new products handler with the given logger
func NewProducts(d *data.Products, logger *log.Logger) *Products {
	return &Products{d, logger}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer this should never happen
// as the router ensures that this is a valid number
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err) // should never happen
	}

	return id
}

// getProduct returns the product from the URL
func getProduct(r *http.Request) data.Product {
	return r.Context().Value(KeyProduct{}).(data.Product)
}
