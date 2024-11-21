package domain

type Cargo struct {
	ID          int
	Weight      int
	Volume      float64
	Description string
}

type CargoRepository interface {
	GetByID(id int) (*Cargo, error)
	Create(cargo *Cargo) error
	Update(cargo *Cargo) error
	Delete(cargo *Cargo) error
}
