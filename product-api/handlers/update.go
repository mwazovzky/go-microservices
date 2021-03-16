// curl -v -X PUT http://localhost:9090/2 -d '{"name": "Espresso", "description": "New taste", "price": 3.20, "sku": "fdj777"}'
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

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
}
