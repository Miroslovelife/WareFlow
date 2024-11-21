package main

import (
	"github.com/Miroslovelife/WareFlow/internal/adapter/grpc_ware_flow"
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/pkg/mongo"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/internal/usecase"
	"github.com/Miroslovelife/WareFlow/pkg/simplex"
)

func main() {
	// === Настройка MongoDB клиента ===
	mongoURI := "mongodb://localhost:27017"
	databaseName := "wareflow"

	mongoClient, err := mongo.NewMongoClient(mongoURI, databaseName)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}
	defer func() {
		if err := mongoClient.Close(); err != nil {
			log.Fatalf("Failed to close MongoDB client: %v", err)
		}
	}()

	warehouseCollection := mongoClient.GetCollection("warehouses")
	cargoCollection := mongoClient.GetCollection("cargos")
	transportCollection := mongoClient.GetCollection("transports")
	locationCollection := mongoClient.GetCollection("locations")
	pathCollection := mongoClient.GetCollection("paths")
	optimizationResultCollection := mongoClient.GetCollection("optimization_results")

	warehouseRepo := repository.NewMongoGenericRepository[models.WareHouse](warehouseCollection)
	cargoRepo := repository.NewMongoGenericRepository[models.Cargo](cargoCollection)
	transportRepo := repository.NewMongoGenericRepository[models.Transport](transportCollection)
	locationRepo := repository.NewMongoGenericRepository[models.Location](locationCollection)
	pathRepo := repository.NewMongoGenericRepository[models.Path](pathCollection)
	optimizationResultRepo := repository.NewMongoGenericRepository[models.OptimizationResult](optimizationResultCollection)

	optimizer := simplex.NewSimplexOptimizer()
	fuelPrice := 1.5

	optimizationUseCase := usecase.NewOptimizationUseCase(
		optimizer,
		fuelPrice,
		cargoRepo,
		transportRepo,
		warehouseRepo,
	)

	// === Инициализация gRPC сервера ===
	grpcServer := grpc_ware_flow.NewWareFlowServiceServer(
		warehouseRepo,
		cargoRepo,
		transportRepo,
		locationRepo,
		pathRepo,
		optimizationResultRepo,
		optimizationUseCase,
	)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterWareFlowServiceServer(server, grpcServer)

	log.Println("WareFlow gRPC server is running on port 50051...")

	// Запуск gRPC-сервера
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
