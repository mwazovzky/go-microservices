package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	protos "github.com/mwazovzky/microservices-introduction/currency/protos/currency"
	"github.com/mwazovzky/microservices-introduction/product-api/data"
	"github.com/mwazovzky/microservices-introduction/product-api/handlers"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Connect to gRPC service
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create gRPC client
	cc := protos.NewCurrencyClient(conn)

	// Create products storage/service
	repository := data.NewRepository()
	db := data.NewProducts(repository, cc)

	productsHandler := handlers.NewProducts(db, logger)

	sm := mux.NewRouter()
	sm.Use(productsHandler.MiddlewareLogRequest)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.Index).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products", productsHandler.Index)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.Show).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.Show)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandler.Create)
	postRouter.Use(productsHandler.MiddlwareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.Update)
	putRouter.Use(productsHandler.MiddlwareValidateProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS middleware
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// https://golang.org/pkg/net/http/#Server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      cors(sm),
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
