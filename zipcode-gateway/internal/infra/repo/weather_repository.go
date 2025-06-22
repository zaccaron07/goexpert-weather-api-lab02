package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	CityName              string  `json:"city"`
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

type WeatherRepository struct {
	WeatherAPIURL string
}

func NewWeatherRepository(apiURL string) *WeatherRepository {
	return &WeatherRepository{WeatherAPIURL: apiURL}
}

func (r *WeatherRepository) GetByZipcode(cep string) (WeatherResponse, error) {
	url := fmt.Sprintf("%s/zipcode/%s/weather", r.WeatherAPIURL, cep)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("service returned status %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	var w WeatherResponse
	if err := json.Unmarshal(body, &w); err != nil {
		return WeatherResponse{}, err
	}
	return w, nil
}
