package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type weatherAPIResponse struct {
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

func (r *WeatherRepository) GetByZipcode(ctx context.Context, cep string) (entity.Weather, error) {
	url := fmt.Sprintf("%s/zipcode/%s/weather", r.WeatherAPIURL, cep)
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entity.Weather{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return entity.Weather{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return entity.Weather{}, fmt.Errorf("%s", string(body))
	}

	var apiResp weatherAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return entity.Weather{}, err
	}
	return entity.Weather{
		CityName:              apiResp.CityName,
		TemperatureCelsius:    apiResp.TemperatureCelsius,
		TemperatureFahrenheit: apiResp.TemperatureFahrenheit,
		TemperatureKelvin:     apiResp.TemperatureKelvin,
	}, nil
}
