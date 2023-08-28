package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-demo/helloworld"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)

	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	name := "kishore"
	var age int32 = 20

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name, Age: age})
	if err != nil {
		log.Fatalf("Failed to call say hello: %v", err)
	}

	fmt.Printf("Response: %v\n", response.Message)
}
