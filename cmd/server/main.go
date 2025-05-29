package main

import (
	"chat-server-service/cmd/server/database"
	"chat-server-service/cmd/server/repository"
	chat "chat-server-service/pkg/chatServer_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = 50052

func main() {

	ctx := context.Background()

	lis, lisErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if lisErr != nil {
		log.Fatalf("failed to listen: %v", lisErr)
	}

	dbPool, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	pingErr := dbPool.Ping(ctx)
	if pingErr != nil {
		log.Fatalf("failed to ping database: %v", pingErr)
	}
	defer dbPool.Close()

	serverRepo := repository.NewChatServerRepository(dbPool)

	server := grpc.NewServer()
	reflection.Register(server)
	chat.RegisterChatServerV1Server(server, serverRepo)

	log.Printf("Server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
