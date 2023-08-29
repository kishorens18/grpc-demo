# grpc-demo

Code Flow:

1.The server initializes, connects to the database, and sets up the gRPC server to listen on a specific port.

2.The client initiates a connection to the server.

3.The client constructs a pb.Profile message with the name "Kiran".

4.The client sends the message to the server's CreateProfile method.

5.The server receives the request and routes it to the CreateProfile method in the controller.

6.The controller maps the request to a models.Profile and calls the corresponding service method.

7.The service method interacts with the MongoDB database to create a profile and returns a models.DBResponse.

8.The controller constructs a pro.ProfileResponse based on the service response and sends it back to the client.

9.The client receives the gRPC response and prints the response data.

This code flow illustrates the client-server interaction for creating user profiles in a bank using gRPC and MongoDB.
