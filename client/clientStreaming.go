package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

func callSayClientStreaming(client pb.GreetingsServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error calling SayClientStreaming: %v", err)
	}
	defer stream.CloseSend()
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("Received response: %v", resp)
}
