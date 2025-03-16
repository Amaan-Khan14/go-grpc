package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

func callSayHello(client pb.GreetingsServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", resp.Message)
}
