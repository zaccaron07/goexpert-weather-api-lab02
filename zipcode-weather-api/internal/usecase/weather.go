package usecase

import "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"

type WeatherInput struct {
	CityName string
}

type WeatherOutput struct {
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}
type GetWeatherUseCase struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewGetWeatherUseCase(weatherRepository entity.WeatherRepositoryInterface) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		WeatherRepository: weatherRepository,
	}
}
func (c *GetWeatherUseCase) Execute(input WeatherInput) (WeatherOutput, error) {
	weather, err := c.WeatherRepository.GetByCityName(input.CityName)
	if err != nil {
		return WeatherOutput{}, err
	}

	weatherOutput := WeatherOutput{
		TemperatureCelsius:    weather.TemperatureCelsius,
		TemperatureFahrenheit: weather.TemperatureFahrenheit,
		TemperatureKelvin:     weather.TemperatureKelvin,
	}

	return weatherOutput, nil
}
