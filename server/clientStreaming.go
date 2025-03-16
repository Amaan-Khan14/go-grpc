package main

import (
	"io"
	"log"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
)

func (s *greetingsServer) SayClientStreaming(stream pb.GreetingsService_SayClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with names: %v", req.Name)

		// Iterate through each name in the received slice
		messages = append(messages, "Hello, "+req.Name+"!")
	}

}
