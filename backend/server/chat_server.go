package server

import (
    "context"
    "log"
    "time"

    // Fix: Import proto package sesuai dengan module di go.mod
    "github.com/00dimas/grpc-chat-app/backend/proto"

    // Import MongoDB driver
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// ChatServer struct implements the proto-defined service
type ChatServer struct {
    proto.UnimplementedChatServiceServer
    MongoClient *mongo.Client
}

// SendMessage handles the gRPC call to send a message
func (s *ChatServer) SendMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
    // Access MongoDB collection
    collection := s.MongoClient.Database("chatdb").Collection("messages")

    // Insert message into the database
    _, err := collection.InsertOne(ctx, bson.M{
        "user_id":    req.UserId,
        "group_id":   req.GroupId,
        "text":       req.Text,
        "timestamp":  time.Now(),
    })
    if err != nil {
        log.Printf("Error inserting message: %v", err)
        return nil, err
    }

    return &proto.MessageResponse{
        Status:    "Message sent",
        MessageId: "1234", // Placeholder, replace with actual MongoDB document ID
    }, nil
}

// GetMessages streams messages for a given group
func (s *ChatServer) GetMessages(req *proto.MessageHistoryRequest, stream proto.ChatService_GetMessagesServer) error {
    // Access MongoDB collection
    collection := s.MongoClient.Database("chatdb").Collection("messages")

    // Query messages by group_id
    cursor, err := collection.Find(context.Background(), bson.M{"group_id": req.GroupId})
    if err != nil {
        log.Printf("Error retrieving messages: %v", err)
        return err
    }
    defer cursor.Close(context.Background())

    // Stream messages to the client
    for cursor.Next(context.Background()) {
        var msg proto.Message
        if err := cursor.Decode(&msg); err != nil {
            log.Printf("Error decoding message: %v", err)
            return err
        }
        if err := stream.Send(&msg); err != nil {
            log.Printf("Error streaming message: %v", err)
            return err
        }
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Cursor error: %v", err)
        return err
    }

    return nil
}
