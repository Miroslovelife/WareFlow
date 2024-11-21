package repository

import (
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/internal/repository"
)

type CargoRepository interface {
	repository.GenericRepository[models.Cargo]
}

type LocationRepository interface {
	repository.GenericRepository[models.Location]
}

type OptimizationService interface {
	CalculateOptimalPath(
		warehouses []models.WareHouse,
		transports []models.Transport,
		cargos []models.Cargo,
	) (*models.OptimizationResult, error)
}

type OptimizationResultRepository interface {
	repository.GenericRepository[models.OptimizationResult]
}

type PathRepository interface {
	repository.GenericRepository[models.Path]
}

type TransportRepository interface {
	repository.GenericRepository[models.Transport]
}

type WareHouseRepository interface {
	repository.GenericRepository[models.WareHouse]
}
