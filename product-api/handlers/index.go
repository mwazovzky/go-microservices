// curl -v  http://localhost:9090
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// Index handles GET requests and returns all current products
func (p *Products) Index(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := data.ToJSON(products, rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}
