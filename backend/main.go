package main

import (
    "context"
    "log"
    "net"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    pb "github.com/00dimas/grpc-chat-app/backend/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type ChatServer struct {
    pb.UnimplementedChatServiceServer
}

var messageCollection *mongo.Collection


func connectDB() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    messageCollection = client.Database("grpc_chat").Collection("messages")
    log.Println("Connected to MongoDB!")
}

func (s *ChatServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
    
    _, err := messageCollection.InsertOne(ctx, bson.M{
        "user_id":   req.UserId,
        "group_id":  req.GroupId,
        "text":      req.Text,
        "timestamp": time.Now().Format(time.RFC3339), 
    })
    if err != nil {
        log.Printf("Failed to save message: %v", err)
        return nil, err
    }

    log.Printf("Received message: UserID=%s, GroupID=%s, Text=%s", req.UserId, req.GroupId, req.Text)

    return &pb.MessageResponse{
        Status:    "Message sent successfully",
        MessageId: "12345", 
    }, nil
}

func (s *ChatServer) GetMessages(req *pb.MessageHistoryRequest, stream pb.ChatService_GetMessagesServer) error {
    cursor, err := messageCollection.Find(context.TODO(), bson.M{"group_id": req.GroupId})
    if err != nil {
        return err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var doc struct {
            UserID    string `bson:"user_id"`
            GroupID   string `bson:"group_id"`
            Text      string `bson:"text"`
            Timestamp string `bson:"timestamp"`
        }
        if err := cursor.Decode(&doc); err != nil {
            return err
        }

        msg := &pb.Message{
            UserId:    doc.UserID,
            GroupId:   doc.GroupID,
            Text:      doc.Text,
            Timestamp: doc.Timestamp,
        }
        if err := stream.Send(msg); err != nil {
            return err
        }
    }

    return nil
}

func main() {
    
    connectDB()

    
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen on port 50051: %v", err)
    }

    grpcServer := grpc.NewServer()

    
    pb.RegisterChatServiceServer(grpcServer, &ChatServer{})

    
    reflection.Register(grpcServer)

    log.Println("gRPC Server is running on port 50051...")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}