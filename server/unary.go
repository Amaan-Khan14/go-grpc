package main

import (
	"context"
	"log"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

var count int

func (s *greetingsServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	count++
	log.Printf("Request: %v", count) ///just for reference
	return &pb.HelloResponse{
		Message: "Hello! from server",
	}, nil
}
