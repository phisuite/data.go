# Phi Suite Data.go

| **Homepage** | [https://phisuite.com][0]        |
| ------------ | -------------------------------- | 
| **GitHub**   | [https://github.com/phisuite][1] |

## Overview

This project contains the Go module to create the **Data API** server & client.  
For more details, see [Phi Suite Data][2].

## Installation

```bash
go get github.com/phisuite/data.go
```

## Creating the server

```go
package main
import "github.com/phisuite/data.go"

type eventAPIServer struct { ... }

func (s *eventAPIServer) Publish(ctx context.Context, req *data.Event) (*data.Event, error) {
    ... 
}

lis, err := net.Listen("tcp", ":50051")
grpcServer := grpc.NewServer()
data.RegisterEventAPIServer(grpcServer, &eventAPIServer{})
grpcServer.Serve(lis)
```
For more details, see [gRPC Basics - Go: Creating the server][10].

## Creating the client

```go
package main
import "github.com/phisuite/data.go"

conn, err := grpc.Dial(*serverAddr)
defer conn.Close()

client := data.NewEventAPIClient(conn)
event := client.Publish(context.Background(), &data.Event{ ... })
```
For more details, see [gRPC Basics - Go: Creating the client][11].

[0]: https://phisuite.com
[1]: https://github.com/phisuite
[2]: https://github.com/phisuite/data
[10]: https://www.grpc.io/docs/tutorials/basic/go/#server
[11]: https://www.grpc.io/docs/tutorials/basic/go/#client
