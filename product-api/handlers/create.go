// curl -v -X POST http://localhost:9090 -d '{"name": "tea", "description": "Nice cup of tea", "price": 0.99, "sku": "xyz987"}'
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	product := getProduct(r)

	data.AddProduct(product)
}
