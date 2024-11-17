package main

import (
	"log"
	"net"

	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("./cfg.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация сервера с конфигурацией
	server, err := InitializeServer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// Настройка gRPC сервера
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWareFlowServiceServer(grpcServer, server)

	// Включаем рефлексию для отладки
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
