package main

import (
	"log"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":50051"
)

func main() {
	clientConn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer clientConn.Close()

	client := pb.NewGreetingsServiceClient(clientConn)
	// callSayHello(client)
	//
	names := &pb.NamesList{
		Name: []string{
			"Alice",
			"Bob",
			"John Doe",
		},
	}
	callSayHelloServerStreaming(client, names)
}
