# GO Microservices

## Swagger code generation

[goswagger.io](https://goswagger.io/)

```
swagger generate client -f ../swagger.yaml -A product-api
```

## Open api documentation generation

```
GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
```

Documenation is available at [docs](http://localhost:9090/docs)

## CORS

[Understanding CORS](https://medium.com/@baphemot/understanding-cors-18ad6b478e2b)
