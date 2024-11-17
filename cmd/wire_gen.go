//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Miroslovelife/WareFlow/internal/config"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/internal/server"
	"github.com/Miroslovelife/WareFlow/internal/usecase"
	"github.com/Miroslovelife/WareFlow/pkg/mongo"
	"github.com/Miroslovelife/WareFlow/pkg/simplex"
	"github.com/google/wire"
)

func InitializeServer(cfg *config.Config) (*server.OptimizationServiceServer, error) {
	wire.Build(
		// Преобразуем конфигурацию
		wire.Value(mongo.MongoConfig{
			URI:    cfg.MongoURI,
			DBName: cfg.DBName,
		}),

		// Fuel price provider
		wire.Value(1.5), // Топливная цена

		// MongoDB client
		mongo.NewMongoClient,

		// Создаем коллекции
		wire.Bind(new(*mongo.Collection), wire.Struct(new(mongo.MongoClient), "client", "dbName")),
		wire.Value("warehouse"),
		wire.Value("cargo"),
		wire.Value("transport"),
		wire.Value("location"),
		wire.Value("path"),
		wire.Value("optimization"),

		// Repositories
		repository.NewMongoWarehouseRepository,
		wire.Bind(new(domain.WareHouseRepository), new(*repository.WarehouseRepositoryMongo)),

		repository.NewMongoCargoRepository,
		wire.Bind(new(domain.CargoRepository), new(*repository.CargoRepositoryMongo)),

		repository.NewTransportRepositoryMongo,
		wire.Bind(new(domain.TransportRepository), new(*repository.TransportRepositoryMongo)),

		repository.NewMongoLocationRepository,
		wire.Bind(new(domain.LocationRepository), new(*repository.LocationRepositoryMongo)),

		repository.NewMongoPathRepository,
		wire.Bind(new(domain.PathRepository), new(*repository.PathRepositoryMongo)),

		repository.NewMongoOptimizationResultRepository,
		wire.Bind(new(domain.OptimizationResultRepository), new(*repository.OptimizationResultRepositoryMongo)),

		// UseCase
		simplex.NewSimplexOptimizer,
		usecase.NewOptimizationUseCase,
		wire.Bind(new(domain.OptimizationService), new(*usecase.OptimizationUseCase)),

		// Server
		server.NewOptimizationServiceServer,
	)

	return &server.OptimizationServiceServer{}, nil
}
