# Currency-Conversion-Service
* This is a REST API layer built to connect to the currency-exchange-service microservice(using gRPC) and get the latest currency rates.

* Protoc Code Gen Tool:
    ```
  export PATH="$PATH:$(go env GOPATH)/bin"
  go get -u github.com/golang/protobuf/protoc-gen-go
    ```
* Generate proto code:
    ````
  protoc -I=$SRC_DIR --go_out=plugins=$DST_DIR $SRC_DIR/currency-exchange.proto
  protoc -I=proto/ --go_out=plugins=grpc:. proto/currency-exchange.proto
    ````