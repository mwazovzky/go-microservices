// curl -v -X DELETE http://localhost:9090/products/2
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
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
