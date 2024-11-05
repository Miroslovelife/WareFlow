package domain

type OptimizationResult struct {
	TransportID   int
	TotalDistance float64
	TotalCost     float64
	Route         []Path
}

type OptimizationResultRepository interface {
	GetByID(id int) (*OptimizationResult, error)
	Create(path *OptimizationResult) error
	Update(path *OptimizationResult) error
	Delete(path *OptimizationResult) error
}
