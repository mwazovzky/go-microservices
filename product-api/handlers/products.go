package handlers

import (
	"context"
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

/*
curl -v  http://localhost:9090
*/
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	list := data.GetProducts()

	err := list.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

/*
curl -v -X POST http://localhost:9090 -d '{"name": "tea", "description": "Nice cup of tea", "price": 0.99, "sku": "xyz987"}'
*/
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := getProduct(r)

	data.AddProduct(product)
}

/*
curl -v -X PUT http://localhost:9090/2 -d '{"name": "Espresso", "description": "New taste", "price": 3.20, "sku": "fdj777"}'
*/
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to parse id", http.StatusBadRequest)
		return
	}

	product := getProduct(r)

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Failed to update product", http.StatusInternalServerError)
		return
	}
}

func getProduct(r *http.Request) *data.Product {
	return r.Context().Value(KeyProduct{}).(*data.Product)
}

type KeyProduct struct{}

func (p Products) MiddlwareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
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
