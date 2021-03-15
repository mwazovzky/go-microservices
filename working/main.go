package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mwazovzky/microservices-introduction/working/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	handlerHello := handlers.NewHello(logger)
	handlerGoodbye := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", handlerHello)
	sm.Handle("/goodbye", handlerGoodbye)

	log.Println("Starting http server at :9090")
	http.ListenAndServe(":9090", sm)
}
