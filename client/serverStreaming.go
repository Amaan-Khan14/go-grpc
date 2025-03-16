package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

func callSayHelloServerStreaming(client pb.GreetingsServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Error calling SayHelloServerStreaming: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		log.Printf("Received response: %s", response.Message)
	}

	log.Println("Streaming completed")
}
