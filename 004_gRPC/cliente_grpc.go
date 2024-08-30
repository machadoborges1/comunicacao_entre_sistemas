//Cliente gRPC: Conecta-se ao servidor, envia uma mensagem e espera por uma resposta.

package main

import (
	"context"
	"log"
	pb "path/to/generated/chat"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Send a message
	err = stream.Send(&pb.ChatMessage{User: "User1", Message: "Hello!"})
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Receive a message
	msg, err := stream.Recv()
	if err != nil {
		log.Fatalf("Error receiving message: %v", err)
	}
	log.Printf("Received message from %s: %s", msg.User, msg.Message)
}
