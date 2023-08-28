package main

import (
	"context"
	"fmt"
	pb "grpc-demo/customer"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type customerServiceServer struct {
	mu        sync.Mutex
	customers map[string]*pb.Customer
	pb.UnimplementedCustomerserviceServer
}

func (s *customerServiceServer) InsertCustomer(ctx context.Context, req *pb.Customer) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var nextAccId int32 = int32(1000 + len(s.customers))
	req.AccId = nextAccId
	s.customers[req.Id] = req

	return &pb.CustomerResponse{Id: req.Id}, nil
}

func (s *customerServiceServer) GetCustomerDetails(ctx context.Context, req *pb.CustomerID) (*pb.CustomerDetails, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	customer, ok := s.customers[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Customer not found")
	}

	return &pb.CustomerDetails{Customer: customer}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen:%v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterCustomerserviceServer(server, &customerServiceServer{
		customers: make(map[string]*pb.Customer),
	})
	fmt.Println("Server listening on :50051")
	if err := server.Serve(lis); err != nil {
		fmt.Printf("Failed to server: %v", err)
	}
}
