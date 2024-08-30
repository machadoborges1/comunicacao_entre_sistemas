//Servidor gRPC: Implementa o método Chat, que recebe mensagens do cliente, imprime no log, e ecoa a mensagem de volta ao cliente.

package main

import (
	"log"
	"net"
	pb "path/to/generated/chat"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message from %s: %s", msg.User, msg.Message)
		// Echo the message back to the client
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
