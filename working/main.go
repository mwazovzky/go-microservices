package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mwazovzky/microservices-introduction/working/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	handlerHello := handlers.NewHello(logger)
	handlerGoodbye := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", handlerHello)
	sm.Handle("/goodbye", handlerGoodbye)
	// https://golang.org/pkg/net/http/#Server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Starting http server at :9090")
	server.ListenAndServe()
}
