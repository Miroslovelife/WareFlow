package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Структуры для работы с геокодированием через Nominatim
type NominatimResponse []struct {
    Lat string `json:"lat"`
    Lon string `json:"lon"`
}

// Функция для геокодирования
func GeocodeAddress(address string) (string, string, error) {
    addressEncoded := url.QueryEscape(address)
    url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", addressEncoded)
    resp, err := http.Get(url)
    if err != nil {
        return "", "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", "", err
    }

    var result NominatimResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return "", "", err
    }

    if len(result) == 0 {
        return "", "", fmt.Errorf("no results found for address %s", address)
    }

    return result[0].Lat, result[0].Lon, nil
}
