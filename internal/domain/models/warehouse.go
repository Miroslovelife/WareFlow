package domain

type WareHouse struct {
	ID       int
	Location *Location
}

type WareHouseRepository interface {
	GetByID(id int) (*WareHouse, error)
	Create(warehouse *WareHouse) error
	Update(warehouse *WareHouse) error
	Delete(id int) error
}
