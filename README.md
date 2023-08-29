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


// Connection between these are : 

1. Client-Server Connection:

Client (client.go): The client-side application initiates communication with the server by creating a gRPC connection using grpc.Dial. It connects to the server running on localhost:5001 and specifies that the connection should be insecure (unencrypted) using grpc.WithInsecure().

Server (server.go): The server-side application creates a gRPC server using grpc.NewServer(). It then registers the implementation of the gRPC service, which is the controllers.RPCServer{} struct, using pro.RegisterProfileServiceServer(s, &controllers.RPCServer{}). The server listens on port 5001 for incoming gRPC requests.

2. Server-Database Connection:

Server (server.go): Within the main function of the server, a connection to the MongoDB database is established using the ConnectDataBase function from the config package. This function returns a MongoDB client instance.

Server (server.go): After connecting to the database, the server initializes the customer service by calling initDatabase(mongoclient). This involves passing the MongoDB client to the initDatabase function.

Services (services.go): The ProfileService struct holds a reference to the MongoDB collection (ProfileCollection) and the context. This context is necessary for MongoDB operations.

Services (services.go): The InitProfileService function initializes the customer service. It does this by creating a ProfileService struct and assigning the provided MongoDB collection and context to it.

3. Client-Server Communication:

Client (client.go): The client application constructs a gRPC request message (pb.Profile) with the desired profile name (e.g., "Kiran"). It then invokes the CreateProfile method on the gRPC client (client.CreateProfile), passing the request message.

Server (controllers.go): The server receives the gRPC request and routes it to the CreateProfile method within the RPCServer implementation. This method corresponds to the CreateProfile RPC defined in the protobuf definition.

Server (controllers.go): Inside the CreateProfile method, the received request is mapped to a models.Profile struct. This struct represents the data that will be stored in the MongoDB database.

Server (services.go): The service method CreateProfile is then called with the mapped models.Profile. The service carries out the following steps:

Sets various properties of the profile (e.g., CreatedAt, UpdatedAt, etc.).
Inserts the profile into the MongoDB collection using InsertOne.
Checks for errors during insertion, such as duplicate emails.
Server (services.go): If the insertion is successful, the service creates a unique index for the email field to prevent duplicate emails in the future.

Server (services.go): Finally, the service queries the database to retrieve the newly created profile and constructs a models.DBResponse based on it.

Server (controllers.go): The CreateProfile method in the controller constructs a pro.ProfileResponse using the data from the models.DBResponse. This response is returned to the gRPC framework, which will send it back to the client.

Client (client.go): The client receives the gRPC response and extracts the response.Name (assuming the response contains the created profile's name). This extracted name is printed to the console.
