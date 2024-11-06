package warehouse

import (
	"math"
)

// Структура для склада с информацией о стоимости, расходе топлива и грузоподъемности
type Warehouse struct {
    Name                  string
    Address               string
    CostPerKm             float64 // Базовая стоимость за километр
    FuelConsumptionPer100Km float64 // Расход топлива на 100 км
    MaxVolume             float64 // Максимальный объем груза (м³)
    MaxWeight             float64 // Максимальный вес груза (кг)
}

// Метод для расчета общей стоимости маршрута с учетом расхода топлива и необходимости дополнительных машин
func (w Warehouse) CalculateTotalCost(distance float64, fuelPrice, cargoVolume, cargoWeight float64) (float64, error) {
    // Проверка, что склад может вместить вес и объем груза
    requiredTrucks := 1

    if cargoVolume > w.MaxVolume || cargoWeight > w.MaxWeight {
        requiredTrucks = int(math.Ceil(max(cargoVolume/w.MaxVolume, cargoWeight/w.MaxWeight)))
    }

    // Базовая стоимость с учетом расстояния и затрат на бензин
    baseCost := (distance / 1000) * w.CostPerKm * float64(requiredTrucks)
    fuelCost := (distance / 1000) * (w.FuelConsumptionPer100Km / 100) * fuelPrice * float64(requiredTrucks)

    return baseCost + fuelCost, nil
}

// Вспомогательная функция для нахождения максимального значения
func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
