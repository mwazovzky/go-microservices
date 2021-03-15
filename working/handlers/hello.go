package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World!")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
		h.logger.Printf("Bad request: %s", err)
		return
	}

	h.logger.Printf("Data: %s", data)
	fmt.Fprintf(rw, "Hello %s!\n", data)
}
