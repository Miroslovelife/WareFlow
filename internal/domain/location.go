package domain

type Location struct {
	ID        int
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
}

type LocationRepository interface {
	GetByID(id int) (Location, error)
	Create(location Location) error
	Update(location Location) error
	Delete(location Location) error
}
