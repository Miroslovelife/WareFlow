package models

type Transport struct {
	ID             int
	Type           string
	CapacityVolume int
	CapacityWeight int
	Expense        float64
}
