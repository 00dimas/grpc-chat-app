package server

import (
    "context"
    "log"
    "time"
    "grpc-chat-app/backend/proto"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type ChatServer struct {
    proto.UnimplementedChatServiceServer
    MongoClient *mongo.Client
}

func (s *ChatServer) SendMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
    collection := s.MongoClient.Database("chatdb").Collection("messages")
    _, err := collection.InsertOne(ctx, bson.M{
        "user_id":    req.UserId,
        "group_id":   req.GroupId,
        "text":       req.Text,
        "timestamp":  time.Now(),
    })
    if err != nil {
        return nil, err
    }
    return &proto.MessageResponse{
        Status:    "Message sent",
        MessageId: "1234", // Placeholder, ganti sesuai ID MongoDB
    }, nil
}

func (s *ChatServer) GetMessages(req *proto.MessageHistoryRequest, stream proto.ChatService_GetMessagesServer) error {
    collection := s.MongoClient.Database("chatdb").Collection("messages")
    cursor, err := collection.Find(context.Background(), bson.M{"group_id": req.GroupId})
    if err != nil {
        return err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var msg proto.Message
        if err := cursor.Decode(&msg); err != nil {
            return err
        }
        if err := stream.Send(&msg); err != nil {
            return err
        }
    }
    return nil
}
