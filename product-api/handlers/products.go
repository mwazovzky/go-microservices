package handlers

import (
	"log"
	"net/http"

	"github.com/mwazovzky/microservices-introduction/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catch all other
	rw.WriteHeader(http.StatusMethodNotAllowed) // 405
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	list := data.GetProducts()

	err := list.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

/*
curl -v -X POST http://localhost:9090 -d '{"name": "tea", "description": "Nice cup of tea", "price": 0.99, "sku": "xyz987"}'
*/
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
