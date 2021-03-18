package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mwazovzky/microservices-introduction/product-images/files"
	"github.com/mwazovzky/microservices-introduction/product-images/handlers"
)

var port string
var logLevel string
var basePath string

func init() {
	godotenv.Load()

	port = os.Getenv("PORT")
	logLevel = os.Getenv("LOG_LEVEL")
	basePath = os.Getenv("BASE_PATH")
}

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create storage
	store, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		logger.Println("Unable to create a storage", err)
		os.Exit(1)
	}
	// create handlers
	fh := handlers.NewFiles(store, logger)

	// create mux/router
	sm := mux.NewRouter()

	// create subrouter to upload files
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	// create subrouter to download files
	// curl -v localhost:9090/images/1/test.png -o test2.png
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))),
	)

	// https://golang.org/pkg/net/http/#Server
	server := &http.Server{
		Addr:         port,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Starting http server at", port)
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
