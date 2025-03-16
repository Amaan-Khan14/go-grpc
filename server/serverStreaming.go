package main

import (
	"fmt"
	"time"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

func (s *greetingsServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetingsService_SayHelloServerStreamingServer) error {
	fmt.Printf("Received request with %s names", req.Name)

	for _, name := range req.Name {
		resp := &pb.HelloResponse{
			Message: "Hello " + name + " from server",
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
