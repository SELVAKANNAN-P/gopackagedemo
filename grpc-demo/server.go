package main

import (
	"context"
	"fmt"
	"net"

	pb "grpc-demo/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HellpoRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello,%s!", req.Name),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed to  listen :%v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	fmt.Println("server listening on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve %v", err)
	}
}
