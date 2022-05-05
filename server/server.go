package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ibonny/golang-tests/protocol"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
