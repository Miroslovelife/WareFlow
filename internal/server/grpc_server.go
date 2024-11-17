package server

import (
	"context"

	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/domain"
)

type OptimizationServiceServer struct {
	pb.UnimplementedWareFlowServiceServer
	Optimization           domain.OptimizationService
	WarehouseRepo          domain.WareHouseRepository
	CargoRepo              domain.CargoRepository
	TransportRepo          domain.TransportRepository
	LocationRepo           domain.LocationRepository           // Добавлено
	PathRepo               domain.PathRepository               // Добавлено
	OptimizationResultRepo domain.OptimizationResultRepository // Добавлено
}

func NewOptimizationServiceServer(
	optimizationService domain.OptimizationService,
	warehouseRepo domain.WareHouseRepository,
	cargoRepo domain.CargoRepository,
	transportRepo domain.TransportRepository,
	locationRepo domain.LocationRepository,
	pathRepo domain.PathRepository,
	optimizationResultRepo domain.OptimizationResultRepository,
) *OptimizationServiceServer {
	return &OptimizationServiceServer{
		Optimization:           optimizationService,
		WarehouseRepo:          warehouseRepo,
		CargoRepo:              cargoRepo,
		TransportRepo:          transportRepo,
		LocationRepo:           locationRepo,
		PathRepo:               pathRepo,
		OptimizationResultRepo: optimizationResultRepo,
	}
}

func (s *OptimizationServiceServer) CalculateOptimalPath(ctx context.Context, req *pb.OptimizationRequest) (*pb.OptimizationResponse, error) {
	// Преобразуем запрос в доменные структуры
	var warehouses []domain.WareHouse
	var transports []domain.Transport
	var cargos []domain.Cargo

	for _, w := range req.Warehouses {
		warehouses = append(warehouses, domain.WareHouse{
			ID: int(w.Id),
			Location: &domain.Location{
				ID:        int(w.Location.Id),
				Name:      w.Location.Name,
				Address:   w.Location.Address,
				Latitude:  w.Location.Latitude,
				Longitude: w.Location.Longitude,
			},
		})
	}

	for _, t := range req.Transports {
		transports = append(transports, domain.Transport{
			ID:             int(t.Id),
			Type:           t.Type,
			CapacityVolume: int(t.CapacityVolume),
			CapacityWeight: int(t.CapacityWeight),
			Expense:        t.Expense,
		})
	}

	for _, c := range req.Cargos {
		cargos = append(cargos, domain.Cargo{
			ID:          int(c.Id),
			Weight:      int(c.Weight),
			Volume:      c.Volume,
			Description: c.Description,
		})
	}

	// Вызываем оптимизатор
	result, err := s.Optimization.CalculateOptimalPath(warehouses, transports, cargos)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	var paths []*pb.Path
	for _, p := range result.Route {
		paths = append(paths, &pb.Path{
			StartLocationId: int32(p.StartLocationID),
			EndLocationId:   int32(p.EndLocationID),
			Distance:        p.Distance,
			Duration:        p.Duration,
			FuelPrice:       p.FuelPrice,
		})
	}

	return &pb.OptimizationResponse{
		TotalDistance: result.TotalDistance,
		TotalCost:     result.TotalCost,
		Routes:        paths,
	}, nil
}
