package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
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
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "product-images",
		Level: hclog.LevelFromString(logLevel),
	})

	// create storage
	store, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		logger.Error("Unable to create a storage", err)
		os.Exit(1)
	}
	// create handlers
	fh := handlers.NewFiles(store, logger)

	// create mux/router
	sm := mux.NewRouter()

	// create subrouter to upload files
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.Upload)
	ph.HandleFunc("/images", fh.UploadMultipart)

	// create subrouter for get requests
	gh := sm.Methods(http.MethodGet).Subrouter()
	// curl -v localhost:9090/images/1
	gh.HandleFunc("/images/{id:[0-9]+}", fh.Index)
	// curl -v localhost:9090/images/1/test.png -o test2.png
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))),
	)

	// CORS middleware
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// https://golang.org/pkg/net/http/#Server
	server := &http.Server{
		Addr:         port,
		Handler:      cors(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Info("Starting http server", "port", port)

		err := server.ListenAndServe()
		if err != nil {
			logger.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until a signal is received.
	sig := <-sigChan
	logger.Info("Recieved terminate signal, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 10 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(ctx)
}
