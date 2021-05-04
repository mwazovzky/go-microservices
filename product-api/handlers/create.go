// curl -v -X POST http://localhost:9090/products -d '{"name": "tea", "description": "Nice cup of tea", "price": 0.99, "sku": "xyz987"}'
package handlers

import (
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	product := getProduct(r)

	p.data.AddProduct(product)

	rw.WriteHeader(http.StatusNoContent)
}
