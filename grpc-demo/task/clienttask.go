package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-demo/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	Id:= 4
	response, err := client.SayHello(context.Background(), &pb.TaskResponse{id: Id})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Message)
}
