package main

import (
	"context"
	"log"
	"time"

	pb "github.com/anujagrawal699/grpc-go/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming RPC")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloClientStreaming RPC: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request to server: %v", err)
		}
		log.Printf("Request sent with name: %v", name)
		time.Sleep(2 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client steaming finished")
	if err != nil {
		log.Fatalf("Error while receiving response from server: %v", err)
	}
	log.Printf("Response from server: %v", res.Messages)
}
