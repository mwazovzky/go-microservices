package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mwazovzky/microservices-introduction/product-api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/", productsHandler)
	// https://golang.org/pkg/net/http/#Server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Starting http server at :9090")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// gracefully shutdown the server allows to complete current request
	sigChan := make(chan os.Signal)
	// broadcast operating system signals to the channel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// wait for the signal
	sig := <-sigChan
	logger.Println("Recieved terminate signal, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(ctx)
}
