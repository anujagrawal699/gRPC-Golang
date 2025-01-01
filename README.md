# gRPC Golang Project

A simple gRPC service implementation in Go demonstrating unary, server streaming, client streaming, and bidirectional streaming communication patterns.

## Prerequisites

- Go >= 1.18
- Protocol Buffers Compiler (protoc)
- gRPC Go plugins

## Installation

```bash
# Clone the repository
git clone https://github.com/anujagrawal699/anujagrawal699-gRPC-Golang.git
cd anujagrawal699-gRPC-Golang

# Install dependencies
go mod tidy

# Generate gRPC code
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## Usage

### Running the Server
```bash
cd server
go run *.go
```

### Running the Client
```bash
cd client
go run *.go
```

## Features

The project demonstrates four gRPC communication patterns:

- **Unary RPC**: Basic request-response model
- **Server Streaming**: Server sends multiple responses to a single client request
- **Client Streaming**: Client sends multiple requests to the server
- **Bidirectional Streaming**: Both client and server exchange multiple messages
