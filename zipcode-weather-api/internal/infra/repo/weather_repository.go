package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherRepository struct {
	APIKey string
}

func NewWeatherRepository(apiKey string) *WeatherRepository {
	return &WeatherRepository{APIKey: apiKey}
}

func (r *WeatherRepository) GetByCityName(ctx context.Context, cityName string) (entity.Weather, error) {
	ctx, span := otel.Tracer("").Start(ctx, "WeatherAPI Lookup")
	span.SetAttributes(attribute.String("city", cityName))
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	encodedCityName := url.QueryEscape(cityName)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", r.APIKey, encodedCityName)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		span.RecordError(err)
		log.Printf("Failed to initialize request: %v", err)
		return entity.Weather{}, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		span.RecordError(err)
		log.Printf("Request failed: %v", err)
		return entity.Weather{}, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		span.RecordError(err)
		log.Printf("Error reading the response: %v", err)
		return entity.Weather{}, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(resp, &weatherResponse)
	if err != nil {
		span.RecordError(err)
		log.Printf("Error parsing response: %v", err)
		return entity.Weather{}, err
	}

	weather := entity.NewWeather(weatherResponse.Current.TempC)
	return *weather, nil
}
