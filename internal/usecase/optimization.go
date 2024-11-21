package usecase

import (
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/internal/repository"
	"github.com/Miroslovelife/WareFlow/pkg/simplex"
)

type OptimizationUseCase struct {
	optimizer     *simplex.SimplexOptimizer
	fuelPrice     float64
	cargoRepo     *repository.MongoGenericRepository[models.Cargo]
	transportRepo *repository.MongoGenericRepository[models.Transport]
	warehouseRepo *repository.MongoGenericRepository[models.WareHouse]
}

func NewOptimizationUseCase(
	optimizer *simplex.SimplexOptimizer,
	fuelPrice float64,
	cargoRepo *repository.MongoGenericRepository[models.Cargo],
	transportRepo *repository.MongoGenericRepository[models.Transport],
	warehouseRepo *repository.MongoGenericRepository[models.WareHouse],
) *OptimizationUseCase {
	return &OptimizationUseCase{
		optimizer:     optimizer,
		fuelPrice:     fuelPrice,
		cargoRepo:     cargoRepo,
		transportRepo: transportRepo,
		warehouseRepo: warehouseRepo,
	}
}

func (u *OptimizationUseCase) CalculateOptimalPath(
	warehouses []models.WareHouse,
	transports []models.Transport,
	cargos []models.Cargo,
) (*models.OptimizationResult, error) {

	paths := []models.Path{
		{StartLocationID: 1, EndLocationID: 2, Distance: 100.0, Duration: 1.5},
		{StartLocationID: 2, EndLocationID: 3, Distance: 150.0, Duration: 2.0},
	}

	numPaths := len(paths)
	numTransports := len(transports)

	coefficients := make([]float64, numPaths*numTransports)
	for i, path := range paths {
		for j, transport := range transports {
			coefficients[i*numTransports+j] = path.Distance * u.fuelPrice / transport.Expense
		}
	}

	var constraints [][]float64
	var bounds []float64

	for _, transport := range transports {
		volumeConstraint := make([]float64, numPaths*numTransports)
		weightConstraint := make([]float64, numPaths*numTransports)
		for p := 0; p < numPaths; p++ {
			volumeConstraint[p*numTransports] = 1
			weightConstraint[p*numTransports] = 1
		}
		constraints = append(constraints, volumeConstraint, weightConstraint)
		bounds = append(bounds, float64(transport.CapacityVolume), float64(transport.CapacityWeight))
	}

	for _, cargo := range cargos {
		cargoVolumeConstraint := make([]float64, numPaths*numTransports)
		cargoWeightConstraint := make([]float64, numPaths*numTransports)
		for p := 0; p < numPaths; p++ {
			cargoVolumeConstraint[p*numTransports] = 1
			cargoWeightConstraint[p*numTransports] = 1
		}
		constraints = append(constraints, cargoVolumeConstraint, cargoWeightConstraint)
		bounds = append(bounds, cargo.Volume, float64(cargo.Weight))
	}

	variableBounds := make([][2]float64, numPaths*numTransports)
	for i := range variableBounds {
		variableBounds[i] = [2]float64{0, 1}
	}

	solution, objective, err := u.optimizer.Minimize(coefficients, constraints, bounds, variableBounds)
	if err != nil {
		return nil, err
	}

	var selectedPaths []models.Path
	var selectedTransportIDs []int
	totalDistance := 0.0

	for i, val := range solution {
		if val > 0.5 {
			pathIndex := i / numTransports
			transportIndex := i % numTransports

			selectedPaths = append(selectedPaths, paths[pathIndex])
			selectedTransportIDs = append(selectedTransportIDs, transports[transportIndex].ID)
			totalDistance += paths[pathIndex].Distance
		}
	}

	return &models.OptimizationResult{
		TransportID:   selectedTransportIDs,
		TotalDistance: totalDistance,
		TotalCost:     objective,
		Route:         selectedPaths,
	}, nil
}
