package repository

import "WareFlow/internal/domain"

type DistanceRepository interface {
	GetDistance(start, end domain.Location) (float64, error)
}
