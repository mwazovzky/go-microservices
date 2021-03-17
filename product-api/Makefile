swagger: 
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	cd swagger-codegen && swagger generate client -f ../swagger.yaml -A product-api