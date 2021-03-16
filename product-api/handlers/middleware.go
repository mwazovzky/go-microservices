package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
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

// MiddlewareLogRequest logs request data and calls next
func (p Products) MiddlewareLogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.logger.Printf("Handle %s Product\n", r.Method)
		next.ServeHTTP(rw, r)
	})
}
