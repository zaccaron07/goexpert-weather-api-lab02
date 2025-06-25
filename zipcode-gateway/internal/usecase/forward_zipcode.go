package usecase

import (
	"context"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/entity"
)

type ForwardZipcodeInput struct {
	Zipcode string
}

type ForwardZipcodeOutput struct {
	CityName              string
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

type WeatherUseCase struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewWeatherUseCase(weatherRepository entity.WeatherRepositoryInterface) *WeatherUseCase {
	return &WeatherUseCase{WeatherRepository: weatherRepository}
}

func (w *WeatherUseCase) ForwardZipcode(ctx context.Context, input ForwardZipcodeInput) (ForwardZipcodeOutput, error) {
	zipCodeEntity, err := entity.NewZipcode(input.Zipcode)
	if err != nil {
		return ForwardZipcodeOutput{}, err
	}

	weather, err := w.WeatherRepository.GetByZipcode(ctx, zipCodeEntity.Cep)
	if err != nil {
		return ForwardZipcodeOutput{}, err
	}

	forwardZipcodeOutput := ForwardZipcodeOutput{
		CityName:              weather.CityName,
		TemperatureCelsius:    weather.TemperatureCelsius,
		TemperatureFahrenheit: weather.TemperatureFahrenheit,
		TemperatureKelvin:     weather.TemperatureKelvin,
	}

	return forwardZipcodeOutput, nil
}
