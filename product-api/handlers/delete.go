// curl -v -X DELETE http://localhost:9090/2
package handlers

import (
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

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
