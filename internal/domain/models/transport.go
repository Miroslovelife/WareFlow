package domain

type Transport struct {
	ID             int
	Type           string
	CapacityVolume int
	CapacityWeight int
	Expense        float64
}

type TransportRepository interface {
	GetByID(id int) (*Transport, error)
	Create(transport *Transport) error
	Update(transport *Transport) error
	Delete(id int) error
}
