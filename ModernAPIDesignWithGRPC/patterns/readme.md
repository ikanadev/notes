### Example 1: Async & Sync REST Fibonacci
Run go server `go run ./cmd/rest-fibonacci-server/main.go`
Run sync client `go run ./cmd/rest-fibonacci-client/rest-fibonacci-client.go -typeOfCall=sync -number=45`
Run async client `go run ./cmd/rest-fibonacci-client/rest-fibonacci-client.go -typeOfCall=async -number=45`

Sync will wait to the numbers to be calculated before returning.
Async will return the current calculated numbers each second until it finishes.
