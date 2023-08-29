# grpc-demo

how the code flow happens:

1.The gRPC server is started using the main function in server/server.go.

2.The server initializes the customer service and associates the RPCServer implementation with the gRPC server.

3.The client initiates a gRPC request by executing the main function in client/client.go.

4.The client establishes a connection with the gRPC server.

5.The client constructs a Customer message and sends it as a request to the server.

6.The server receives the request and routes it to the CreateCustomer method in controllers/controllers.go.

7.The controller maps the request to a customer data structure and calls the corresponding service method.

8.The service method (services/customer.go) interacts with the MongoDB database to create a customer and returns the response.

9.The controller converts the service response to a gRPC response and sends it back to the client.

10.The client receives the gRPC response, which includes the customer ID if the customer creation was successful.

Overall, the client sends a request to the server, the server processes the request through the controller and service layers, interacts with the database if needed, and sends a response back to the client. This code flow enables communication between the client and server for customer creation operations.
