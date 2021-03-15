package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		p.updateProduct(rw, r)
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

/*
curl -v -X PUT http://localhost:9090/2 -d '{"name": "Espresso", "description": "New taste", "price": 3.20, "sku": "fdj777"}'
*/
func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Product")

	id, err := parseID(r.URL.Path)
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	product := &data.Product{}

	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Failed to update product", http.StatusInternalServerError)
		return
	}
}

func parseID(path string) (int, error) {
	var err error
	reg := regexp.MustCompile(`/([0-9]+)`)
	group := reg.FindAllStringSubmatch(path, -1)

	if len(group) != 1 {
		return 0, fmt.Errorf("Invalid URI")
	}

	if len(group[0]) != 2 {
		return 0, fmt.Errorf("Invalid URI")
	}

	idString := group[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}

	return id, nil
}
