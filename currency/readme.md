# User Guide

## GRPC

### Useful links

[Protocol Buffers - Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
[Protocol Buffers - Encoding](https://developers.google.com/protocol-buffers/docs/encoding)
[Protocol Buffers - Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
[GRPC quick start](https://grpc.io/docs/languages/go/quickstart/)
[grpcCurl](https://github.com/fullstorydev/grpcurl)

### Install grpc

```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Generate service

```
make protoc
```

### Make grpc request

need to register reflection for the service to be able get service info
can be turned of for production environment

Get service info using reflection

```
$ grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection

$ grpcurl --plaintext localhost:9092 list Currency
Currency.GetRate

$ grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );

$ grpcurl --plaintext localhost:9092 describe .RateRequest
RateRequest is a message:
message RateRequest {
  string Base = 1;
  string Destination = 2;
}

$ grpcurl --plaintext localhost:9092 describe .RateResponse
RateResponse is a message:
message RateResponse {
  float Rate = 1;
}
```

Call servive methods

```
$ grpcurl --plaintext -d '{"Base": "EUR", "Destination": "USD"}' localhost:9092 Currency.GetRate
{
  "Rate": 1.2082
}
```

## Testing

### Run tests

```
go test -v ./data
```
