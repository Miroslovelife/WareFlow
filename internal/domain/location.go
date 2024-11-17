package domain

type Location struct {
	ID        int
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
}

type LocationRepository interface {
	GetByID(id int) (Location, error) // Возвращает Location, а не указатель
	Create(location *Location) error
	Update(location *Location) error
	Delete(id int) error
}
