package usecase

import (
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

func (w *WeatherUseCase) ForwardZipcode(input ForwardZipcodeInput) (ForwardZipcodeOutput, error) {
	zipCodeEntity, err := entity.NewZipcode(input.Zipcode)
	if err != nil {
		return ForwardZipcodeOutput{}, err
	}

	weatherDto, err := w.WeatherRepository.GetByZipcode(zipCodeEntity.Cep)

	if err != nil {
		return ForwardZipcodeOutput{}, err
	}

	forwardZipcodeOutput := ForwardZipcodeOutput{
		CityName:              weatherDto.CityName,
		TemperatureCelsius:    weatherDto.TemperatureCelsius,
		TemperatureFahrenheit: weatherDto.TemperatureFahrenheit,
		TemperatureKelvin:     weatherDto.TemperatureKelvin,
	}

	return forwardZipcodeOutput, nil
}
