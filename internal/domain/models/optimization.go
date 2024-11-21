package models

// OptimizationResult содержит результаты оптимизации
type OptimizationResult struct {
	TransportID   []int
	TotalDistance float64
	TotalCost     float64
	Route         []Path
}
