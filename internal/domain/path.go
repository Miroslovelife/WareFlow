package domain

type Path struct {
	StartLocationID int
	EndLocationID   int
	Distance        float64
	Duration        float64
	FuelPrice       float64
}

type PathRepository interface {
	GetByID(id int) (*Path, error)
	Create(path *Path) error
	Update(path *Path) error
	Delete(path *Path) error
}
