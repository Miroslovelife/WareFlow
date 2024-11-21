package grpc

import (
	"context"
	pb "github.com/Miroslovelife/WareFlow/github.com/miroslav/WareFlowV2/proto"
	"github.com/Miroslovelife/WareFlow/internal/domain/models"
	"github.com/Miroslovelife/WareFlow/internal/usecase"
)

type OptimizationServiceServer struct {
	pb.UnimplementedWareFlowServiceServer
	usecase *usecase.OptimizationUseCase
}

func NewOptimizationServiceServer(usecase *usecase.OptimizationUseCase) *OptimizationServiceServer {
	return &OptimizationServiceServer{
		usecase: usecase,
	}
}

func (s *OptimizationServiceServer) CalculateOptimalPath(
	ctx context.Context,
	req *pb.OptimizationRequest,
) (*pb.OptimizationResponse, error) {
	warehouses := mapWarehousesFromProto(req.Warehouses)
	transports := mapTransportsFromProto(req.Transports)
	cargos := mapCargosFromProto(req.Cargos)

	result, err := s.usecase.CalculateOptimalPath(warehouses, transports, cargos)
	if err != nil {
		return nil, err
	}

	response := mapOptimizationResultToProto(result)
	return response, nil
}

func mapWarehousesFromProto(protoWarehouses []*pb.Warehouse) []models.WareHouse {
	var warehouses []models.WareHouse
	for _, w := range protoWarehouses {
		warehouses = append(warehouses, models.WareHouse{
			ID: int(w.Id),
			Location: &models.Location{
				ID:        int(w.Location.Id),
				Name:      w.Location.Name,
				Address:   w.Location.Address,
				Latitude:  w.Location.Latitude,
				Longitude: w.Location.Longitude,
			},
		})
	}
	return warehouses
}

func mapTransportsFromProto(protoTransports []*pb.Transport) []models.Transport {
	var transports []models.Transport
	for _, t := range protoTransports {
		transports = append(transports, models.Transport{
			ID:             int(t.Id),
			Type:           t.Type,
			CapacityVolume: int(t.CapacityVolume),
			CapacityWeight: int(t.CapacityWeight),
			Expense:        t.Expense,
		})
	}
	return transports
}

func mapCargosFromProto(protoCargos []*pb.Cargo) []models.Cargo {
	var cargos []models.Cargo
	for _, c := range protoCargos {
		cargos = append(cargos, models.Cargo{
			ID:          int(c.Id),
			Weight:      int(c.Weight),
			Volume:      c.Volume,
			Description: c.Description,
		})
	}
	return cargos
}

func mapOptimizationResultToProto(result *models.OptimizationResult) *pb.OptimizationResponse {
	var routes []*pb.Path
	for _, p := range result.Route {
		routes = append(routes, &pb.Path{
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
		Routes:        routes,
	}
}
