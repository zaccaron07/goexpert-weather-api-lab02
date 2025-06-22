package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/infra/repo"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/usecase"
)

type WeatherResponse struct {
	City                  string  `json:"city"`
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

type WebWeatherHandler struct {
	ZipcodeRepository entity.ZipcodeRepositoryInterface
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewWebWeatherHandler(apiKey string) *WebWeatherHandler {
	return &WebWeatherHandler{
		ZipcodeRepository: repo.NewZipcodeRepository(),
		WeatherRepository: repo.NewWeatherRepository(apiKey),
	}
}

func (h *WebWeatherHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")

	getZipcodeInput := usecase.ZipcodeInput{
		CEP: zipcode,
	}

	getZipcodeUseCase := usecase.NewGetZipcodeUseCase(h.ZipcodeRepository)
	getZipcodeOutput, err := getZipcodeUseCase.Execute(getZipcodeInput)

	if err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, fmt.Sprintf("error fetching location: %v", err), http.StatusInternalServerError)
		return
	}
	if getZipcodeOutput.Localidade == "" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	getWeatherInput := usecase.WeatherInput{
		CityName: getZipcodeOutput.Localidade,
	}
	getWeatherUseCase := usecase.NewGetWeatherUseCase(h.WeatherRepository)
	getWeatherOutput, err := getWeatherUseCase.Execute(getWeatherInput)

	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching weather: %v", err), http.StatusInternalServerError)
		return
	}

	weatherResponse := WeatherResponse{
		City:                  getZipcodeOutput.Localidade,
		TemperatureCelsius:    getWeatherOutput.TemperatureCelsius,
		TemperatureFahrenheit: getWeatherOutput.TemperatureFahrenheit,
		TemperatureKelvin:     getWeatherOutput.TemperatureKelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResponse)
}
