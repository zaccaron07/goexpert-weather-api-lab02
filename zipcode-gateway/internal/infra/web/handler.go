package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/repo"
	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/usecase"
)

type WeatherResponse struct {
	City                  string  `json:"city"`
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

type WebWeatherHandler struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewWebWeatherHandler(weatherAPIUrl string) *WebWeatherHandler {
	return &WebWeatherHandler{
		WeatherRepository: repo.NewWeatherRepository(weatherAPIUrl),
	}
}

func (h *WebWeatherHandler) ZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Cep string `json:"cep"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if reqBody.Cep == "" {
		http.Error(w, "zipcode is required", http.StatusBadRequest)
		return
	}

	useCaseInput := usecase.ForwardZipcodeInput{
		Zipcode: reqBody.Cep,
	}

	weatherUseCase := usecase.NewWeatherUseCase(h.WeatherRepository)
	forwardZipcodeOutput, err := weatherUseCase.ForwardZipcode(r.Context(), useCaseInput)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "invalid zipcode" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if strings.TrimSpace(err.Error()) == "can not find zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := WeatherResponse{
		City:                  forwardZipcodeOutput.CityName,
		TemperatureCelsius:    forwardZipcodeOutput.TemperatureCelsius,
		TemperatureFahrenheit: forwardZipcodeOutput.TemperatureFahrenheit,
		TemperatureKelvin:     forwardZipcodeOutput.TemperatureKelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
