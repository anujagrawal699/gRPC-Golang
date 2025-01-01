package main

import (
	"context"

	pb "github.com/anujagrawal699/grpc-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from gRPC server!",
	}, nil
}
