package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Структуры для ответа от OSRM
type RouteResponse struct {
    Routes []struct {
        Legs []struct {
            Distance float64 `json:"distance"`
            Duration float64 `json:"duration"`
        } `json:"legs"`
    } `json:"routes"`
}

// Функция для получения расстояния до магазина
func GetDistanceToStore(lat, lon string, storeLat, storeLon string) (float64, float64, error) {
    url := fmt.Sprintf("http://localhost:5000/route/v1/driving/%s,%s;%s,%s?overview=false", lon, lat, storeLon, storeLat)
    resp, err := http.Get(url)
    if err != nil {
        return 0, 0, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, 0, err
    }

    var result RouteResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return 0, 0, err
    }

    if len(result.Routes) == 0 || len(result.Routes[0].Legs) == 0 {
        return 0, 0, fmt.Errorf("no routes found")
    }

    distance := result.Routes[0].Legs[0].Distance // в метрах
    duration := result.Routes[0].Legs[0].Duration // в секундах

    return distance, duration, nil
}
