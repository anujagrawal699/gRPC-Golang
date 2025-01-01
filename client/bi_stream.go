package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/anujagrawal699/grpc-go/proto"
)

func callSayHelloBiDiStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bi-Directional Streaming Started")
	stream, err := client.SayHelloBiDiStreaming(context.Background())
	if err != nil {
		log.Fatalf("Failed to call SayHelloBiDiStreaming: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive a message: %v", err)
			}
			log.Printf("Response from server: %v", message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
		time.Sleep(5 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi-Directional Streaming Ended")
}
