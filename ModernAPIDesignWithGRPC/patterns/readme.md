### Example 1: Async & Sync REST Fibonacci
Run go server `go run ./cmd/rest-fibonacci-server/main.go`
Run sync client `go run ./cmd/rest-fibonacci-client/rest-fibonacci-client.go -typeOfCall=sync -number=45`
Run async client `go run ./cmd/rest-fibonacci-client/rest-fibonacci-client.go -typeOfCall=async -number=45`

Sync will wait to the numbers to be calculated before returning.
Async will return the current calculated numbers each second until it finishes.


### Example 2: Async (server stream) & Sync gRPC Fibonacci
Generate the proto code: `protoc --go_out=. --go-grpc_out=. proto/*.proto`
Run server `go run ./cmd/grpc-fibonacci-server/main.go`
Run sync client `go run ./cmd/grpc-fibonacci-client/main.go -typeOfCall=sync -number=45`
Run async client `go run ./cmd/grpc-fibonacci-client/main.go -typeOfCall=async -number=45`

### Example 3: Client Stream gRPC: Find average number
Client sends a stream of numbers and server returns the average when it finishes
Generate the proto code: `protoc --go_out=. --go-grpc_out=. proto/*.proto`
Run server `go run ./cmd/grpc-average-server/main.go`
Run client `go run ./cmd/grpc-average-client/main.go -numbers=1,2,3,4`

### Example 4: Bi-directional streaming
Client sends a stream of numbers and server returns a stream of the max number
Generate the proto code: `protoc --go_out=. --go-grpc_out=. proto/*.proto`
Run server `go run ./cmd/grpc-max-server/main.go`
Run client `go run ./cmd/grpc-max-client/main.go -numbers=5,2,3,4,5,6,7,8,9,10,11,13,15,25`
