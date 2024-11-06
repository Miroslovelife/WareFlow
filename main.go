package main

import (
	"fmt"
	"os/exec"
	"routing/geo"
	"routing/route"
	"routing/warehouse"
	"strconv"
)

func main() {
    // Актуальная цена на бензин (рублей за литр)
    fuelPrice := 55.0

    // Объем и вес груза (например, объем в кубических метрах и вес в кг)
    cargoVolume := 20.0 // кубометров
    cargoWeight := 1500.0 // кг

    // Определяем склады с машинами, расходом топлива и грузоподъемностью
    warehouses := []warehouse.Warehouse{
        {Name: "Склад 1", Address: "Великие Луки", CostPerKm: 30.0, FuelConsumptionPer100Km: 12.0, MaxVolume: 15.0, MaxWeight: 1000.0},
        {Name: "Склад 2", Address: "Санкт-Петербург", CostPerKm: 25.0, FuelConsumptionPer100Km: 10.0, MaxVolume: 20.0, MaxWeight: 1500.0},
        {Name: "Склад 3", Address: "Москва", CostPerKm: 20.0, FuelConsumptionPer100Km: 8.0, MaxVolume: 25.0, MaxWeight: 2000.0},
    }

    storeAddress := "Смоленск"
    
    // Получаем координаты магазина
    storeLat, storeLon, err := geo.GeocodeAddress(storeAddress)
    if err != nil {
        fmt.Printf("Ошибка геокодирования магазина: %v\n", err)
        return
    }

    // Создаем слайс для передачи параметров в Python-скрипт
    var args []string

    // Проверяем стоимость маршрута от каждого склада до магазина
    for _, wh := range warehouses {
        lat, lon, err := geo.GeocodeAddress(wh.Address)
        if err != nil {
            fmt.Printf("Ошибка геокодирования склада %s: %v\n", wh.Name, err)
            continue
        }

        distance, duration, err := route.GetDistanceToStore(lat, lon, storeLat, storeLon)
        if err != nil {
            fmt.Printf("Ошибка получения маршрута для склада %s: %v\n", wh.Name, err)
            continue
        }

        // Рассчитываем стоимость с учетом объема и веса груза
        cost, err := wh.CalculateTotalCost(distance, fuelPrice, cargoVolume, cargoWeight)
        if err != nil {
            fmt.Printf("Ошибка расчета стоимости для склада %s: %v\n", wh.Name, err)
            continue
        }

        // Добавляем параметры маршрута в слайс аргументов для Python
        args = append(args, strconv.FormatFloat(cost, 'f', 2, 64))
        args = append(args, strconv.FormatFloat(duration/3600, 'f', 2, 64))
        args = append(args, strconv.FormatFloat(distance/1000, 'f', 2, 64))

        fmt.Printf("Склад: %s, Расстояние: %.2f км, Время в пути: %.2f часов, Стоимость: %.2f\n", wh.Name, distance/1000, duration/3600, cost)
    }

    // Устанавливаем ограничения
    args = append(args, "11", "1000") // Максимально допустимое время и расстояние


    // Запускаем Python-скрипт с аргументами
    cmd := exec.Command("python", append([]string{"lp_solver.py"}, args...)...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Ошибка при выполнении Python скрипта: %v\n", err)
        return
    }

    // Вывод результата Python-скрипта
    fmt.Println(string(output))
}
