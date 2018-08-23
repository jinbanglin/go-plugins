# JSON-RPC 2.0 Codec

## Usage

Import the codec and set within the client/server
```go
package main

import (
    "github.com/jinbanglin/go-micro"
    "github.com/jinbanglin/go-micro/client"
    "github.com/jinbanglin/go-micro/server"
    "github.com/jinbanglin/go-plugins/codec/jsonrpc2"
)

func main() {
    client := client.NewClient(
        client.Codec("application/json", jsonrpc2.NewCodec),
        client.ContentType("application/json"),
    )

    server := server.NewServer(
        server.Codec("application/json", jsonrpc2.NewCodec),
    )

    service := micro.NewService(
        micro.Client(client),
        micro.Server(server),
    )

    // ...
}
```

