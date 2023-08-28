package main

import (
	"context"
	"fmt"
	pb "grpc-demo/bank"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerServiceServer struct {
	mu          sync.Mutex
	customers   map[string]*pb.Customer
	mongoClient *mongo.Client
}

func (s *customerServiceServer) InsertCustomer(ctx context.Context, req *pb.Customer) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a document.
	document := bson.D{
		{Key: "id", Value: req.Id},
		{Key: "name", Value: req.Name},
		{Key: "accid", Value: req.Accid},
		{Key: "balance", Value: req.Balance},
	}

	// Insert the document into the database.
	_, err := s.mongoClient.Database("mydb").Collection("customers").InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}

	return &pb.CustomerResponse{Id: req.Id}, nil
}

func (s *customerServiceServer) GetCustomerDetails(ctx context.Context, req *pb.CustomerID) (*pb.CustomerDetails, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the customer exists.
	customer, ok := s.customers[req.Id]
	if !ok {
		return nil, fmt.Errorf("Customer not found")
	}

	return &pb.CustomerDetails{Customer: customer}, nil
}
