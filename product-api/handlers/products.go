package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

// GetProducts handles GET requests and returns all current products
// curl -v  http://localhost:9090
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := data.ToJSON(products, rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

// AddProduct handles POST requests to add new products
// curl -v -X POST http://localhost:9090 -d '{"name": "tea", "description": "Nice cup of tea", "price": 0.99, "sku": "xyz987"}'
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := getProduct(r)

	data.AddProduct(product)
}

/*
curl -v -X PUT http://localhost:9090/2 -d '{"name": "Espresso", "description": "New taste", "price": 3.20, "sku": "fdj777"}'
*/
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	product := getProduct(r)
	product.ID = getProductID(r)

	err := data.UpdateProduct(product)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Failed to update product", http.StatusInternalServerError)
		return
	}
}

// Delete handles DELETE requests and removes items
// curl -v -X DELETE http://localhost:9090/2
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.logger.Println("[ERROR] Product not found")

		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.logger.Println("[ERROR] Failed to delete product")

		http.Error(rw, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
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

type KeyProduct struct{}

func (p Products) MiddlwareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}

		err := data.FromJSON(&product, r.Body)
		if err != nil {
			p.logger.Println("[ERROR] Unable to decode request body")
			http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			p.logger.Println("[ERROR] Validation error")
			http.Error(
				rw,
				fmt.Sprintf("Validation error: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

func (p Products) MiddlewareLogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.logger.Printf("Handle %s Product\n", r.Method)
		next.ServeHTTP(rw, r)
	})
}
