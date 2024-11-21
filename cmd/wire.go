//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Miroslovelife/WareFlow/internal/adapter/grpc_ware_flow"
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/internal/usecase"
	"github.com/Miroslovelife/WareFlow/pkg/mongo"
	"github.com/Miroslovelife/WareFlow/pkg/simplex"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	// MongoDB клиент
	mongo.NewMongoClient,

	// Репозитории
	repository.NewMongoGenericRepository[models.Cargo],
	repository.NewMongoGenericRepository[models.Location],
	repository.NewMongoGenericRepository[models.Transport],
	repository.NewMongoGenericRepository[models.Path],
	repository.NewMongoGenericRepository[models.WareHouse],
	repository.NewMongoGenericRepository[models.OptimizationResult],

	// Бизнес-логика
	simplex.NewSimplexOptimizer,
	usecase.NewOptimizationUseCase,

	// gRPC сервер
	grpc_ware_flow.NewWareFlowServiceServer,
)

// InitializeServer создаёт сервер приложения
func InitializeServer(uri string, databaseName string) (*grpc_ware_flow.WareFlowServiceServer, func(), error) {
	wire.Build(WireSet)
	return nil, nil, nil
}
