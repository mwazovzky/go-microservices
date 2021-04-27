module github.com/mwazovzky/microservices-introduction/currency

go 1.16

require (
	github.com/mwazovzky/microservices-introduction/currency/protos/currency v0.0.0
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	google.golang.org/grpc v1.36.1 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/mwazovzky/microservices-introduction/currency/protos/currency v0.0.0 => ./protos/currency
