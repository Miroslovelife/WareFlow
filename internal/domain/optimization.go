package domain

type OptimizationService interface {
	CalculateOptimalPath(
		warehouses []WareHouse,
		transports []Transport,
		cargos []Cargo,
	) (*OptimizationResult, error)
}

// OptimizationResult содержит результаты оптимизации
type OptimizationResult struct {
	TransportID   []int
	TotalDistance float64
	TotalCost     float64
	Route         []Path
}

type OptimizationResultRepository interface {
	GetByID(id int) (*OptimizationResult, error)
	Create(result *OptimizationResult) error
	Update(result *OptimizationResult) error
	Delete(result *OptimizationResult) error // Ожидается указатель на OptimizationResult
}
