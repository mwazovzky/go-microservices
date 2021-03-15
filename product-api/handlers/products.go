package handlers

import (
	"encoding/json"
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
	productIndex := data.GetProducts()
	// Encoder allows to write directly to io and avoid keeping large amount of data in memory
	// It is also faster then Marshal
	err := json.NewEncoder(rw).Encode(productIndex)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}
