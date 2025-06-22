package entity

import "github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-gateway/internal/infra/repo"

type WeatherRepositoryInterface interface {
	GetByZipcode(string) (repo.WeatherResponse, error)
}
