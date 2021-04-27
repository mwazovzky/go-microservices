// need to implement CurrencyServer interface generated by protoc
package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	protos "github.com/mwazovzky/microservices-introduction/currency/protos/currency"
)

type CurrencyServer struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrencyServer(l hclog.Logger) *CurrencyServer {
	return &CurrencyServer{log: l}
}

func (cs *CurrencyServer) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	cs.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
