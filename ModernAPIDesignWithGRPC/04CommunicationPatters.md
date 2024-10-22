### Unary RPC

It refers to a simple client-server interaction whre the client sends a single request to the server and receives a single response in return.
This is similar to traditional HTTP request-response.

#### Server-side streaming

The server sends multiple responses to a single client  request over a  single long-lived connection.
Common applications include:
- Real time updates: live scores, stock market, online gaming, video streaming, chat, etc.
- Log streaming: monitoring logs
- Data Feeds: social media feed, news updates, IoT sensor data.
- Batch processing progress: progress updates during batch processing tasks, so the client can monitor the status and progress of such tasks.

The proto definition for the server side streaming RPC is as follows:

```proto
service MyService {
    rpc ServerStreamingMethod(MyRequest) returns (stream MyResponse);
}
```

#### Client-side streaming



#### Bi-directional streaming
