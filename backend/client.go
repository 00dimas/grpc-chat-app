package main

import (
	"context"
	"log"
	"time"

	pb "github.com/00dimas/grpc-chat-app/backend/proto"
	"google.golang.org/grpc"
)

func main() {
	// Hubungkan ke server gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// Kirim pesan
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Uji SendMessage
	resp, err := client.SendMessage(ctx, &pb.MessageRequest{
		UserId:  "user123",
		GroupId: "group123",
		Text:    "Hello, this is a test message!",
	})
	if err != nil {
		log.Fatalf("Error while calling SendMessage: %v", err)
	}
	log.Printf("Response from SendMessage: Status=%s, MessageID=%s", resp.Status, resp.MessageId)

	// Uji GetMessages
	stream, err := client.GetMessages(ctx, &pb.MessageHistoryRequest{
		GroupId: "group123",
	})
	if err != nil {
		log.Fatalf("Error while calling GetMessages: %v", err)
	}

	log.Println("Messages received:")
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("Message from %s: %s", msg.UserId, msg.Text)
	}
}
