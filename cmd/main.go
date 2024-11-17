package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/internal/server"
	"github.com/Miroslovelife/WareFlow/pkg/simplex"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Подключение к MongoDB
	mongoClient, err := connectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	// Создание репозиториев
	pathRepo := repository.NewMongoPathRepository(mongoClient, "WareFlowDB", "Paths")
	cargoRepo := repository.NewMongoCargoRepository(mongoClient, "WareFlowDB", "Cargos")
	locationRepo := repository.NewMongoLocationRepository(mongoClient, "WareFlowDB", "Locations")
	transportRepo := repository.NewTransportRepositoryMongo(mongoClient.Database("WareFlowDB"))

	// Создание бизнес-логики
	optimizer := simplex.NewSimplexOptimizer() // Ваш оптимизатор
	optimizationUseCase := server.NewOptimizationService(pathRepo, cargoRepo, locationRepo, transportRepo, optimizer)

	// Настройка gRPC сервера
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Регистрация сервиса с внедрением зависимостей
	pb.RegisterOptimizationServiceServer(grpcServer, optimizationUseCase)

	// Добавление поддержки Reflection API
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")

	// Запуск gRPC сервера
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Подключение к MongoDB
func connectToMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Проверка соединения
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
