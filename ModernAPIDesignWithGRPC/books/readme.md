####  Run the grpc server
- Setup the database credentials at `configs/grpc-books-server.yaml`
- Run `go run ./cmd/grpc-books-server/main.go -configFile=configs/grpc-books-server.yaml`


####  Run the grpc client
- Run `go run ./cmd/grpc-books-client/main.go -configFile=configs/grpc-books-client.yaml`
