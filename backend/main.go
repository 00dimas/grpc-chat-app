package main

import (
    "log"
    "net"
    "grpc-chat-app/backend/proto"
    "grpc-chat-app/backend/server"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "google.golang.org/grpc"
    "github.com/00dimas/grpc-chat-app/backend/proto"
    "github.com/00dimas/grpc-chat-app/backend/server"
)

func main() {
    mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen on port 50051: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterChatServiceServer(grpcServer, &server.ChatServer{MongoClient: mongoClient})

    log.Println("gRPC server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
