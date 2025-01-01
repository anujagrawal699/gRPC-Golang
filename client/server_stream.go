package main

import (
	"context"
	"io"
	"log"

	pb "github.com/anujagrawal699/grpc-go/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Server Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Failed to call SayHelloServerStreaming: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}
		log.Printf("Response from server: %v", message.Message)
	}
	log.Printf("Server Streaming Ended")

}
