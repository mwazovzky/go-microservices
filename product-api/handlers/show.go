// curl -v  http://localhost:9090/products/1
package handlers

import (
	"context"
	"net/http"

	protos "github.com/mwazovzky/microservices-introduction/currency/protos/currency"
	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

// swagger:route GET /products/{id} products showProduct
// Return a list of products from the database
// responses:
//	200: productsResponse

// Show handles GET requests and returns specified products
func (p *Products) Show(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	product, err := data.FindProduct(id)

	if err == data.ErrProductNotFound {
		p.logger.Println("[ERROR] Product not found", err)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.logger.Println("[ERROR] Failed to fetch product", err)
		http.Error(rw, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	res, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.logger.Println("[ERROR] Failed to get exchange rate", err)
		http.Error(rw, "Failed to get exchange rate", http.StatusInternalServerError)
		return
	}

	// convert product price
	product.Price = product.Price * res.Rate

	rw.Header().Add("Content-Type", "application/json")
	err = data.ToJSON(product, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.logger.Println("[ERROR] serializing product", err)
		http.Error(rw, "Failed to serialize product", http.StatusInternalServerError)
		return
	}
}
