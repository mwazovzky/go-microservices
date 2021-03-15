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

	// catch all other
	rw.WriteHeader(http.StatusMethodNotAllowed) // 405
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	list := data.GetProducts()
	err := list.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}
