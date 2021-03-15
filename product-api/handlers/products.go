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
	encodedData, err := json.Marshal(productIndex)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
		return
	}

	rw.Write(encodedData)
}
