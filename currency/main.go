package main

import (
	"net"
	"os"

	hclog "github.com/hashicorp/go-hclog"
	protos "github.com/mwazovzky/microservices-introduction/currency/protos/currency"
	"github.com/mwazovzky/microservices-introduction/currency/protos/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
