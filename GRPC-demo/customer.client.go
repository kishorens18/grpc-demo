package main

import (
	"context"
	"fmt"
	pb "grpc-demo/customer"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// client:=pb.NewGreeterClient(conn)
	client := pb.NewCustomerserviceClient(conn)

	//Add a customer

	customer := &pb.Customer{
		Id:      "1",
		Name:    "Kiran",
		AccId:   0,
		Balance: 10000,
		BankId:  1,
	}
	addResp, err := client.InsertCustomer(context.Background(), customer)

	if err != nil {
		log.Fatalf("Failed to add customer: %v", err)
	}
	fmt.Printf("Added customer with ID: %s\n", addResp.Id)

	// get customer details

	customerID := &pb.CustomerID{Id: "1"}
	customerDetails, err := client.GetCustomerDetails(context.Background(), customerID)
	if err != nil {
		log.Fatalf("Failed to get customer details: %v", err)
	}
	fmt.Println("Customer details:")
	fmt.Printf("ID: %s,Name:%s,AccountId:%v,Balance:%v,BankId:%v\n", customerDetails.Customer.Id, customerDetails.Customer.Name, customerDetails.Customer.AccId, customerDetails.Customer.Balance, customerDetails.Customer.BankId)
}
