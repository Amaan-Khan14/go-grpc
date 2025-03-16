package main

import (
	"log"
	"net"

	pb "github.com/Amaan-Khan14/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type greetingsServer struct {
	pb.GreetingsServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	// Initialize the grpc server
	grpcServer := grpc.NewServer()

	// Register the greetings service server
	pb.RegisterGreetingsServiceServer(grpcServer, &greetingsServer{})
	log.Printf("Server started on port: %v", listener.Addr())

	// Start serving the grpc server (basically it means app.listen() in Node.js)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve on port %s: %v", port, err)
	}
}
