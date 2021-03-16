// curl -v  http://localhost:9090/products
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// Index handles GET requests and returns all current products
func (p *Products) Index(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := data.ToJSON(products, rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}
