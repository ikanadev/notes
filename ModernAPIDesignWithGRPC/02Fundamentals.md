Protocol Buffers (protobuf), is a serialization format and language-agnostic interface description language, offers a versatile way to structure data.
Usually the steps to use it includes:
1. Create a `.proto` file
2. Generate the language specific implementation using `protoc` (the CLI tool)
3. Use the generated code in your project, to serialize and/or deserialize data

Example: A `BookInfo` proto message:

```proto
message BookInfo {
    string isbn = 1; 
    string publisher = 2;
}
```
Each field consist of type name and tag (Ex. type = string, name = isbn, tag = 1).
The serialized data is in a binary format, which consist of three parts for each field: tag, length, and value concatenated together.
Tags are positive integers and can range from `1` to `2^29-1` (536870911).

GRPC uses HTTP/2 as the transport protocol, this has the following advantages over HTTP/1.1:
1. Multiplexing (multiple requests with a single TCP connection)
2. Header compression
3. Flow control: ensures that data is sent at an optimal pace, preventing congestion.
4. Server push: allows the server to push resources to the client (not common in gRPC)

