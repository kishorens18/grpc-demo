package main

import (
	"context"
	"fmt"
	"net"

	hw "grpc-demo/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	hw.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hw.HelloRequest) (*hw.HelloResponse, error) {
	return &hw.HelloResponse{
		Message: fmt.Sprintf("Hello, %s %v", req.Name, req.Age),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Printf("Fail to listen : %v", err)
		return
	}
	s := grpc.NewServer()
	hw.RegisterGreeterServer(s, &server{})

	fmt.Println("Server listening on :50052")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve :%v", err)
	}
}
