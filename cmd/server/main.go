package main

import (
	chat "chat-server-service/pkg/chatServer_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = 50052

type chatServer struct {
	chat.UnimplementedChatServerV1Server
}

func (server *chatServer) Create(ctx context.Context, request *chat.CreateRequest) (*chat.CreateResponse, error) {
	log.Printf("Received: %v", request.GetUsernames())

	retVal := &chat.CreateResponse{
		Id: 123,
	}

	return retVal, nil
}

func main() {
	lis, lisErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if lisErr != nil {
		log.Fatalf("failed to listen: %v", lisErr)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	chat.RegisterChatServerV1Server(server, &chatServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
