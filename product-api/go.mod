module github.com/mwazovzky/microservices-introduction/product-api

go 1.16

require (
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mwazovzky/microservices-introduction/currency/protos/currency v0.0.0
	google.golang.org/grpc v1.37.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/mwazovzky/microservices-introduction/currency/protos/currency v0.0.0 => ../currency/protos/currency
