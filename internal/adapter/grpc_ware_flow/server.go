package grpc_ware_flow

import (
	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/internal/usecase"
)

type WareFlowServiceServer struct {
	pb.UnimplementedWareFlowServiceServer
	warehouseRepo          *repository.MongoGenericRepository[models.WareHouse]
	cargoRepo              *repository.MongoGenericRepository[models.Cargo]
	transportRepo          *repository.MongoGenericRepository[models.Transport]
	locationRepo           *repository.MongoGenericRepository[models.Location]
	pathRepo               *repository.MongoGenericRepository[models.Path]
	optimizationResultRepo *repository.MongoGenericRepository[models.OptimizationResult]
	useCase                *usecase.OptimizationUseCase
}

func NewWareFlowServiceServer(
	warehouseRepo *repository.MongoGenericRepository[models.WareHouse],
	cargoRepo *repository.MongoGenericRepository[models.Cargo],
	transportRepo *repository.MongoGenericRepository[models.Transport],
	locationRepo *repository.MongoGenericRepository[models.Location],
	pathRepo *repository.MongoGenericRepository[models.Path],
	optimizationResultRepo *repository.MongoGenericRepository[models.OptimizationResult],
	useCase *usecase.OptimizationUseCase,
) *WareFlowServiceServer {
	return &WareFlowServiceServer{
		warehouseRepo:          warehouseRepo,
		cargoRepo:              cargoRepo,
		transportRepo:          transportRepo,
		locationRepo:           locationRepo,
		pathRepo:               pathRepo,
		optimizationResultRepo: optimizationResultRepo,
		useCase:                useCase,
	}
}
