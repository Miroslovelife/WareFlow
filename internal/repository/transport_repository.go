package repository

import "WareFlow/internal/domain"

type TransportRepository interface {
	GetByID(id int) (*domain.Transport, error)
	Create(transport *domain.Transport) error
	Update(transport *domain.Transport) error
	Delete(id int) error
}
