package main

import (
	"context"
	"fmt"
	pb "grpc-demo/bank"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// Create a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client.
	client := pb.NewCustomerserviceClient(conn)

	// Create a customer.
	customer := &pb.Customer{
		Id:      "1",
		Name:    "Kiran",
		Accid:   0,
		Balance: 0,
	}

	// Insert the customer.
	resp, err := client.InsertCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Failed to insert customer: %v", err)
	}

	// Print the response.
	fmt.Println(resp)

	// Get the customer details.
	customerDetails, err := client.GetCustomerDetails(context.Background(), &pb.CustomerID{Id: "1"})
	if err != nil {
		log.Fatalf("Failed to get customer details: %v", err)
	}

	// Print the customer details.
	fmt.Println(customerDetails)
}
