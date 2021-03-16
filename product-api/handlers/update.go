// curl -v -X PUT http://localhost:9090/products/2 -d '{"name": "Espresso", "description": "New taste", "price": 3.20, "sku": "fdj777"}'
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests and updates specified items
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
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

	rw.WriteHeader(http.StatusNoContent)
}
